package handler

import (
	"net/http"

	"github.com/andresxlp/go-echo-postgres-hexa/internal/domain/entities"
	"github.com/andresxlp/go-echo-postgres-hexa/internal/domain/ports/repo"
	"github.com/labstack/echo/v4"
)

type UserHandlers interface {
	GetAllUsersHandler(c echo.Context) error
}

type userRouter struct {
	repository repo.UserRepositoryInterface
}

func NewUserHandler(repository repo.UserRepositoryInterface) UserHandlers {
	return &userRouter{repository}
}

func (ur *userRouter) GetAllUsersHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, entities.Message{
		Message: "Success",
		Data:    "Todos los users",
	})
}
