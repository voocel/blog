package http

import (
	"blog/config"
	"blog/internal/http/handler"
	"blog/internal/http/middleware"
	"blog/internal/http/router"
	"blog/internal/repository/mysql"
	"blog/internal/usecase"
	"blog/internal/usecase/repo"
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
	var dbRepo mysql.Repo
	dbRepo, err = mysql.New()
	if err != nil {
		panic(err)
	}
	g := gin.New()
	gin.SetMode(gin.DebugMode)

	userUseCase := usecase.NewUserUseCase(repo.NewUserRepo(dbRepo.GetDbW()))

	g.Use(
		gin.Logger(),
		gin.Recovery(),
		middleware.Logger,
		middleware.CorsMiddleware(),
		middleware.JWTMiddleware(userUseCase),
	)
	g.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, "404 not found!")
	})

	g.GET("/ping", handler.Ping())
	g.StaticFS("/static", gin.Dir("static", false))
	g.StaticFile("/favicon.ico", "./static/favicon.ico")
	s.routerLoad(g, router.GetRouters(dbRepo.GetDbW())...)

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
