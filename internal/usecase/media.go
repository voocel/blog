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
	mediaType := getMediaType(file.Header.Get("Content-Type"))
	uploadPath := config.Conf.App.UploadPath
	if uploadPath == "" {
		uploadPath = "uploads"
	}

	fullUploadPath := filepath.Join("static", uploadPath)
	if err := os.MkdirAll(fullUploadPath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create upload directory: %w", err)
	}

	// Generate unique filename: UUID + original extension
	ext := filepath.Ext(file.Filename)
	uniqueFilename := uuid.New().String() + ext
	savePath := filepath.Join(fullUploadPath, uniqueFilename)

	src, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer src.Close()

	dst, err := os.Create(savePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dst.Close()

	if _, err := dst.ReadFrom(src); err != nil {
		return nil, fmt.Errorf("failed to save file: %w", err)
	}

	url := fmt.Sprintf("%s/static/%s/%s", baseURL, uploadPath, uniqueFilename)

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
		// If database save fails, delete uploaded file
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
	media, err := uc.mediaRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if err := uc.mediaRepo.Delete(ctx, id); err != nil {
		return err
	}

	// Delete file from disk
	if media.Path != "" {
		if err := os.Remove(media.Path); err != nil {
			// If file doesn't exist or deletion fails, only log but don't return error
			// Because database record has been deleted
			fmt.Printf("Warning: failed to delete file %s: %v\n", media.Path, err)
		}
	}

	return nil
}

// getMediaType determines file type based on MIME type
func getMediaType(contentType string) string {
	if strings.HasPrefix(contentType, "image/") {
		return "image"
	}
	if strings.HasPrefix(contentType, "video/") {
		return "video"
	}
	return "document"
}

// getFileExtension gets file extension
func getFileExtension(filename string) string {
	return strings.ToLower(filepath.Ext(filename))
}
