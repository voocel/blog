package http

import (
	"blog/config"
	"blog/internal/http/middleware"
	"blog/internal/http/router"
	"blog/internal/repository/postgres"
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	srv http.Server
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() {
	// 初始化数据库
	dbRepo, err := postgres.New()
	if err != nil {
		panic(err)
	}

	// 创建 Gin 引擎
	g := gin.New()
	gin.SetMode(config.Conf.Mode)

	// 全局中间件
	g.Use(
		gin.Recovery(),
		middleware.RequestLogger(), // 使用自定义请求日志中间件
		middleware.CorsMiddleware(),
	)

	// 404 处理
	g.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "404 not found"})
	})

	// 静态文件服务
	g.StaticFS("/static", gin.Dir("static", false))
	g.StaticFile("/favicon.ico", "./static/favicon.ico")

	// 配置路由
	router.SetupRoutes(g, dbRepo.GetDbW())

	// 启动 HTTP 服务器
	s.srv = http.Server{
		Addr:    config.Conf.Http.Addr,
		Handler: g,
	}

	go func() {
		if err = s.srv.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
			panic(err)
		}
	}()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
