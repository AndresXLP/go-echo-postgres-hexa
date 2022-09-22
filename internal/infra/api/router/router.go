package router

import (
	"github.com/andresxlp/go-echo-postgres-hexa/internal/infra/api/handler"
	"github.com/andresxlp/go-echo-postgres-hexa/internal/infra/api/router/groups"
	"github.com/labstack/echo/v4"
)

type Router struct {
	server      *echo.Echo
	userHandler groups.UserRoutes
}

func StartRouter(server *echo.Echo, userHandler groups.UserRoutes) *Router {
	return &Router{
		server,
		userHandler,
	}
}

func (rtr *Router) Init() {
	basePath := rtr.server.Group("/api/microblog")
	basePath.GET("/health", handler.HealtCheck)
	rtr.userHandler.Resource(basePath)
}
