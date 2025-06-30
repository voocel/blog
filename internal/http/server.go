package http

import (
	"blog/config"
	"blog/internal/http/handler"
	"blog/internal/http/middleware"
	"blog/internal/http/router"
	"blog/internal/repository/postgres"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	srv http.Server
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() {
	var err error
	var dbRepo postgres.Repo
	dbRepo, err = postgres.New()
	if err != nil {
		panic(err)
	}
	
	g := gin.New()
	gin.SetMode(gin.DebugMode)

	g.Use(
		gin.Logger(),
		gin.Recovery(),
		middleware.Logger,
		middleware.CorsMiddleware(),
	)
	
	g.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, "404 not found!")
	})
	
	g.GET("/ping", handler.Ping())
	g.StaticFS("/static", gin.Dir("static", false))
	g.StaticFile("/favicon.ico", "./static/favicon.ico")
	
	// 加载新的API路由
	for _, loadRouter := range router.GetNewRouters(dbRepo.GetDbW()) {
		loadRouter.Load(g)
	}

	srv := http.Server{
		Addr:    config.Conf.Http.Addr,
		Handler: g,
	}
	
	go func() {
		if err = srv.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
			panic(err)
		}
	}()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
