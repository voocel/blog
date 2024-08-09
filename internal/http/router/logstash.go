package router

import (
	"blog/internal/http/handler"
	"github.com/gin-gonic/gin"
)

type logstashRouter struct {
	h *handler.LogstashHandler
}

func newLogstashRouter(h *handler.LogstashHandler) *logstashRouter {
	return &logstashRouter{h: h}
}

func (r *logstashRouter) Load(g *gin.Engine) {
	g.GET("/api/logstash/list", r.h.List)
}
