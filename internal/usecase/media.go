package usecase

import (
	"blog/config"
	"blog/internal/entity"
	"blog/pkg/log"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
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

func (uc *MediaUseCase) Upload(ctx context.Context, file *multipart.FileHeader, baseURL string, uploadType string) (*entity.MediaResponse, error) {
	detectedMime, ext, err := validateUpload(file, uploadType)
	if err != nil {
		return nil, err
	}
	mediaType := getMediaType(detectedMime)

	// Determine upload directory based on type
	var uploadPath string
	if uploadType == "avatar" {
		uploadPath = "avatar"
	} else {
		uploadPath = config.Conf.App.UploadPath
		if uploadPath == "" {
			uploadPath = "uploads"
		}
	}

	fullUploadPath := filepath.Join("static", uploadPath)
	if err := os.MkdirAll(fullUploadPath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create upload directory: %w", err)
	}

	// Generate unique filename: UUID + original extension
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
		MimeType: detectedMime,
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
			log.Warnf("failed to delete file %s: %v", media.Path, err)
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

const (
	maxAvatarSize    int64 = 2 * 1024 * 1024
	maxPostMediaSize int64 = 50 * 1024 * 1024
	maxSniffBytes          = 512
)

var (
	avatarAllowedTypes = map[string]map[string]struct{}{
		"image/jpeg": {".jpg": {}, ".jpeg": {}},
		"image/png":  {".png": {}},
		"image/gif":  {".gif": {}},
		"image/webp": {".webp": {}},
	}
	postAllowedTypes = map[string]map[string]struct{}{
		"image/jpeg":    {".jpg": {}, ".jpeg": {}},
		"image/png":     {".png": {}},
		"image/gif":     {".gif": {}},
		"image/webp":    {".webp": {}},
		"video/mp4":     {".mp4": {}},
		"video/webm":    {".webm": {}},
		"video/quicktime": {".mov": {}},
	}
)

func validateUpload(file *multipart.FileHeader, uploadType string) (string, string, error) {
	if file == nil {
		return "", "", fmt.Errorf("%w: missing file", ErrInvalidArgument)
	}

	maxSize := maxPostMediaSize
	allowed := postAllowedTypes
	if uploadType == "avatar" {
		maxSize = maxAvatarSize
		allowed = avatarAllowedTypes
	}

	if file.Size <= 0 {
		return "", "", fmt.Errorf("%w: empty file", ErrInvalidArgument)
	}
	if file.Size > maxSize {
		return "", "", fmt.Errorf("%w: file too large", ErrInvalidArgument)
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext == "" {
		return "", "", fmt.Errorf("%w: missing file extension", ErrInvalidArgument)
	}

	src, err := file.Open()
	if err != nil {
		return "", "", fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer src.Close()

	buf := make([]byte, maxSniffBytes)
	n, readErr := src.Read(buf)
	if readErr != nil && readErr != io.EOF {
		return "", "", fmt.Errorf("failed to read uploaded file: %w", readErr)
	}
	detectedMime := http.DetectContentType(buf[:n])

	allowedExts, ok := allowed[detectedMime]
	if !ok {
		return "", "", fmt.Errorf("%w: unsupported file type", ErrInvalidArgument)
	}
	if _, ok := allowedExts[ext]; !ok {
		return "", "", fmt.Errorf("%w: file extension mismatch", ErrInvalidArgument)
	}

	return detectedMime, ext, nil
}
