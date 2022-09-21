package providers

import (
	"github.com/andresxlp/go-echo-postgres-hexa/internal/infra/api/router"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

var Container *dig.Container

func BuildContainer() *dig.Container {
	Container = dig.New()

	_ = Container.Provide(func() *echo.Echo {
		return echo.New()
	})

	_ = Container.Provide(router.StartRouter)
	return Container
}
