package router

import (
	"github.com/andresxlp/go-echo-postgres-hexa/internal/infra/api/handler"
	"github.com/labstack/echo/v4"
)

type Router struct {
	server *echo.Echo
}

func StartRouter(server *echo.Echo) *Router {
	return &Router{server}
}

func (rtr *Router) Init() {
	basePath := rtr.server.Group("/api/microblog")
	basePath.GET("/health", handler.HealtCheck)
}
