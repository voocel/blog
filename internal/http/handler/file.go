package handler

import (
	"blog/config"
	"blog/internal/entity"
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type FileHandler struct {
}

func NewFileHandler() *FileHandler {
	return &FileHandler{}
}

// UploadFile 上传文件
func (h *FileHandler) UploadFile(c *gin.Context) {
	// 获取上传的文件
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "文件上传失败: "+err.Error()))
		return
	}
	defer file.Close()

	// 获取可选的路径参数
	uploadPath := c.PostForm("path")
	if uploadPath == "" {
		uploadPath = "uploads" // 默认路径
	}

	// 验证文件类型
	if !isAllowedImageType(header.Filename) {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "不支持的文件类型，仅支持: "+config.Conf.App.ImageAllowExt))
		return
	}

	// 验证文件大小（MB）
	maxSize := int64(config.Conf.App.ImageMaxSize) * 1024 * 1024
	if header.Size > maxSize {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, fmt.Sprintf("文件大小超过限制，最大允许 %dMB", config.Conf.App.ImageMaxSize)))
		return
	}

	// 创建上传目录（static/{uploadPath}/）
	uploadDir := filepath.Join(config.Conf.App.StaticRootPath, uploadPath)
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, entity.NewErrorResponse(500, "创建上传目录失败"))
		return
	}

	// 生成唯一文件名
	uniqueFilename := generateUniqueFilename(header.Filename)
	filePath := filepath.Join(uploadDir, uniqueFilename)

	// 保存文件
	dst, err := os.Create(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.NewErrorResponse(500, "创建文件失败"))
		return
	}
	defer dst.Close()

	// 复制文件内容
	if _, err := io.Copy(dst, file); err != nil {
		c.JSON(http.StatusInternalServerError, entity.NewErrorResponse(500, "保存文件失败"))
		return
	}

	// 构造文件访问URL
	fileURL := fmt.Sprintf("/static/%s/%s", uploadPath, uniqueFilename)

	// 构造相对路径
	relativePath := filepath.Join(uploadPath, uniqueFilename)

	// 获取正确的MimeType
	mimeType := getMimeType(header.Filename)
	if mimeType == "application/octet-stream" && header.Header.Get("Content-Type") != "" {
		mimeType = header.Header.Get("Content-Type")
	}

	// 构造响应
	response := entity.FileResponse{
		ID:           time.Now().UnixNano(),
		Filename:     uniqueFilename,
		OriginalName: header.Filename,
		MimeType:     mimeType,
		Size:         header.Size,
		URL:          fileURL,
		Path:         relativePath,
		Type:         "file",
		CreatedAt:    time.Now().Format(time.RFC3339),
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse(response, "上传成功"))
}

// GetFiles 获取文件列表
func (h *FileHandler) GetFiles(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	path := c.Query("path")
	fileType := c.Query("type")

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	// 默认路径
	if path == "" {
		path = "uploads"
	}

	// 读取上传目录
	uploadDir := filepath.Join(config.Conf.App.StaticRootPath, path)

	var files []entity.FileResponse

	// 遍历目录
	if entries, err := os.ReadDir(uploadDir); err == nil {
		for _, entry := range entries {
			if fileType == "file" && entry.IsDir() {
				continue
			}
			if fileType == "folder" && !entry.IsDir() {
				continue
			}

			info, err := entry.Info()
			if err != nil {
				continue
			}

			var fileTypeStr string
			var fileURL string
			var relativePath string

			if entry.IsDir() {
				fileTypeStr = "folder"
				fileURL = ""
				relativePath = filepath.Join(path, entry.Name())
			} else {
				fileTypeStr = "file"
				fileURL = fmt.Sprintf("/static/%s/%s", path, entry.Name())
				relativePath = filepath.Join(path, entry.Name())
			}

			file := entity.FileResponse{
				ID:           info.ModTime().UnixNano(),
				Filename:     entry.Name(),
				OriginalName: entry.Name(),
				MimeType:     getMimeType(entry.Name()),
				Size:         info.Size(),
				URL:          fileURL,
				Path:         relativePath,
				Type:         fileTypeStr,
				CreatedAt:    info.ModTime().Format(time.RFC3339),
			}
			files = append(files, file)
		}
	}

	// 简单分页处理
	total := len(files)
	start := (page - 1) * pageSize
	end := start + pageSize
	if start > total {
		files = []entity.FileResponse{}
	} else if end > total {
		files = files[start:]
	} else {
		files = files[start:end]
	}

	paginatedData := entity.NewPaginatedResponse(files, total, page, pageSize)
	c.JSON(http.StatusOK, entity.NewSuccessResponse(paginatedData, "获取成功"))
}

// DeleteFile 删除文件
func (h *FileHandler) DeleteFile(c *gin.Context) {
	idStr := c.Param("id")
	if _, err := strconv.ParseInt(idStr, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "文件ID格式错误"))
		return
	}

	// TODO: 根据ID删除具体文件
	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "删除成功"))
}

// CreateFolder 创建文件夹
func (h *FileHandler) CreateFolder(c *gin.Context) {
	var req entity.CreateFolderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "请求参数错误"))
		return
	}

	// 如果路径为空，默认为uploads
	if req.Path == "" {
		req.Path = "uploads"
	}

	folderPath := filepath.Join(config.Conf.App.StaticRootPath, req.Path, req.Name)
	if err := os.MkdirAll(folderPath, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, entity.NewErrorResponse(500, "创建文件夹失败"))
		return
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "创建成功"))
}

// isAllowedImageType 检查文件类型是否允许
func isAllowedImageType(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	allowedExts := strings.Split(config.Conf.App.ImageAllowExt, ",")

	for _, allowedExt := range allowedExts {
		if strings.TrimSpace(allowedExt) == ext {
			return true
		}
	}
	return false
}

// generateUniqueFilename 生成唯一文件名
func generateUniqueFilename(originalFilename string) string {
	ext := filepath.Ext(originalFilename)
	nameWithoutExt := strings.TrimSuffix(originalFilename, ext)

	// 使用时间戳和文件名的MD5生成唯一名称
	hash := md5.Sum([]byte(fmt.Sprintf("%s_%d", nameWithoutExt, time.Now().UnixNano())))
	uniqueName := fmt.Sprintf("%x", hash)

	return uniqueName + ext
}

// getMimeType 根据文件扩展名获取MIME类型
func getMimeType(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	case ".webp":
		return "image/webp"
	default:
		return "application/octet-stream"
	}
}
