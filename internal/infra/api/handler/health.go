package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Health struct {
	Code    int    `json:"status"`
	Message string `json:"message"`
}

func HealtCheck(ctx echo.Context) error {
	response := &Health{
		Code:    http.StatusOK,
		Message: "Active!",
	}
	return ctx.JSON(http.StatusOK, response)
}
