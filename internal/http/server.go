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
	dbRepo, err := postgres.New()
	if err != nil {
		panic(err)
	}

	g := gin.New()
	gin.SetMode(config.Conf.Mode)

	container := router.NewContainer(dbRepo.GetDbW())
	g.Use(
		gin.Recovery(),         // Panic recovery
		middleware.RequestID(), // Request tracing
		middleware.EventLogger(container.SystemEventRepo), // Event logging
		middleware.RequestLogger(),                        // Request logging
		middleware.CorsMiddleware(),                       // CORS
	)

	g.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "404 not found"})
	})

	// Static file service
	g.StaticFS("/static", gin.Dir("static", false))
	g.StaticFile("/favicon.ico", "./static/favicon.ico")

	router.SetupRoutes(g, container)

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
