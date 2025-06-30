package handler

import (
	"blog/internal/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FileHandler struct {
}

func NewFileHandler() *FileHandler {
	return &FileHandler{}
}

// UploadFile 上传文件
func (h *FileHandler) UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "文件上传失败"))
		return
	}
	defer file.Close()

	// todo
	response := entity.FileResponse{
		ID:           "1",
		Filename:     header.Filename,
		OriginalName: header.Filename,
		MimeType:     header.Header.Get("Content-Type"),
		Size:         header.Size,
		URL:          "/uploads/" + header.Filename,
		Path:         "/uploads/" + header.Filename,
		CreatedAt:    "2024-01-01T00:00:00Z",
	}

	c.JSON(http.StatusOK, entity.NewSuccessResponse(response, "上传成功"))
}

// GetFiles 获取文件列表
func (h *FileHandler) GetFiles(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	// path := c.Query("path")
	// fileType := c.Query("type")

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	// todo
	var files []entity.FileResponse

	paginatedData := entity.NewPaginatedResponse(files, 0, page, pageSize)
	c.JSON(http.StatusOK, entity.NewSuccessResponse(paginatedData, "获取成功"))
}

// DeleteFile 删除文件
func (h *FileHandler) DeleteFile(c *gin.Context) {
	idStr := c.Param("id")
	if _, err := strconv.ParseInt(idStr, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "文件ID格式错误"))
		return
	}

	// todo
	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "删除成功"))
}

// CreateFolder 创建文件夹
func (h *FileHandler) CreateFolder(c *gin.Context) {
	var req entity.CreateFolderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, entity.NewErrorResponse(400, "请求参数错误"))
		return
	}

	// todo
	c.JSON(http.StatusOK, entity.NewSuccessResponse[any](nil, "创建成功"))
} 