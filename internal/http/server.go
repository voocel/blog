package http

import (
	"blog/config"
	"blog/internal/http/handler"
	"blog/internal/http/middleware"
	"blog/internal/http/router"
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

func (s *Server) routerLoad(g *gin.Engine, rs ...router.Router) {
	for _, r := range rs {
		r.Load(g)
	}
}

func (s *Server) Run() {
	var err error
	g := gin.New()
	gin.SetMode(gin.DebugMode)

	g.Use(
		gin.Logger(),
		gin.Recovery(),
		middleware.CorsMiddleware(),
	)
	g.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, "404 not found!")
	})

	g.GET("/ping", handler.Ping())
	g.StaticFS("/static", gin.Dir("static", false))
	g.StaticFile("/favicon.ico", "./static/favicon.ico")
	s.routerLoad(g, router.GetRouters()...)

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
