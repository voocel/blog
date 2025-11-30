package usecase

import (
	"blog/config"
	"blog/internal/entity"
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

type MediaUseCase struct {
	mediaRepo MediaRepo
}

func NewMediaUseCase(mediaRepo MediaRepo) *MediaUseCase {
	return &MediaUseCase{mediaRepo: mediaRepo}
}

func (uc *MediaUseCase) Upload(ctx context.Context, file *multipart.FileHeader, baseURL string) (*entity.MediaResponse, error) {
	// 判断文件类型
	mediaType := getMediaType(file.Header.Get("Content-Type"))

	// 获取上传路径配置
	uploadPath := config.Conf.App.UploadPath
	if uploadPath == "" {
		uploadPath = "uploads"
	}

	// 完整的上传目录路径
	fullUploadPath := filepath.Join("static", uploadPath)

	// 确保上传目录存在
	if err := os.MkdirAll(fullUploadPath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create upload directory: %w", err)
	}

	// 生成唯一文件名：UUID + 原始扩展名
	ext := filepath.Ext(file.Filename)
	uniqueFilename := uuid.New().String() + ext

	// 文件保存路径
	savePath := filepath.Join(fullUploadPath, uniqueFilename)

	// 打开上传的文件
	src, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer src.Close()

	// 创建目标文件
	dst, err := os.Create(savePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dst.Close()

	// 复制文件内容
	if _, err := dst.ReadFrom(src); err != nil {
		return nil, fmt.Errorf("failed to save file: %w", err)
	}

	// 构建访问 URL
	url := fmt.Sprintf("%s/static/%s/%s", baseURL, uploadPath, uniqueFilename)

	// 创建数据库记录
	media := &entity.Media{
		URL:      url,
		Name:     file.Filename,
		Type:     mediaType,
		Size:     file.Size,
		MimeType: file.Header.Get("Content-Type"),
		Path:     savePath,
		Date:     time.Now().Format(time.RFC3339),
	}

	if err := uc.mediaRepo.Create(ctx, media); err != nil {
		// 如果数据库保存失败，删除已上传的文件
		os.Remove(savePath)
		return nil, err
	}

	return &entity.MediaResponse{
		ID:   media.ID,
		URL:  media.URL,
		Name: media.Name,
		Type: media.Type,
		Date: media.Date,
	}, nil
}

func (uc *MediaUseCase) List(ctx context.Context) ([]entity.MediaResponse, error) {
	mediaList, err := uc.mediaRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]entity.MediaResponse, len(mediaList))
	for i, m := range mediaList {
		responses[i] = entity.MediaResponse{
			ID:   m.ID,
			URL:  m.URL,
			Name: m.Name,
			Type: m.Type,
			Date: m.Date,
		}
	}

	return responses, nil
}

func (uc *MediaUseCase) Delete(ctx context.Context, id string) error {
	// 先获取文件信息
	media, err := uc.mediaRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// 从数据库删除记录
	if err := uc.mediaRepo.Delete(ctx, id); err != nil {
		return err
	}

	// 删除磁盘上的文件
	if media.Path != "" {
		if err := os.Remove(media.Path); err != nil {
			// 如果文件不存在或删除失败，只记录但不返回错误
			// 因为数据库记录已经删除了
			fmt.Printf("Warning: failed to delete file %s: %v\n", media.Path, err)
		}
	}

	return nil
}

// 根据 MIME 类型判断文件类型
func getMediaType(contentType string) string {
	if strings.HasPrefix(contentType, "image/") {
		return "image"
	}
	if strings.HasPrefix(contentType, "video/") {
		return "video"
	}
	return "document"
}

// 获取文件扩展名
func getFileExtension(filename string) string {
	return strings.ToLower(filepath.Ext(filename))
}
