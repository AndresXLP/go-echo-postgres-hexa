package groups

import (
	"github.com/andresxlp/go-echo-postgres-hexa/internal/infra/api/handler"
	"github.com/labstack/echo/v4"
)

type (
	UserRoutes interface {
		Resource(c *echo.Group)
	}

	userHandlers struct {
		userHands handler.UserHandlers
	}
)

func NewUserRoutes(userHands handler.UserHandlers) UserRoutes {
	return &userHandlers{userHands}
}

func (routes userHandlers) Resource(e *echo.Group) {
	e.GET("/user", routes.userHands.GetAllUsersHandler)
}
