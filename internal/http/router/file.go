package router

import (
	"blog/internal/http/handler"
	"blog/internal/http/middleware"

	"github.com/gin-gonic/gin"
)

type fileRouter struct {
	fileHandler *handler.FileHandler
}

func newFileRouter(fileHandler *handler.FileHandler) Router {
	return &fileRouter{
		fileHandler: fileHandler,
	}
}

func (r *fileRouter) Load(engine *gin.Engine) {
	fileGroup := engine.Group("/api/files")
	fileGroup.Use(middleware.AuthMiddleware())
	{
		fileGroup.POST("/upload", r.fileHandler.UploadFile)
		fileGroup.GET("", r.fileHandler.GetFiles)
		fileGroup.DELETE("/:id", r.fileHandler.DeleteFile)
		fileGroup.POST("/folder", r.fileHandler.CreateFolder)
	}
} 