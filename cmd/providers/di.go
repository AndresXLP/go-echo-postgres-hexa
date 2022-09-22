package providers

import (
	"github.com/andresxlp/go-echo-postgres-hexa/internal/infra/adapters/postgresql"
	"github.com/andresxlp/go-echo-postgres-hexa/internal/infra/adapters/postgresql/repo"
	"github.com/andresxlp/go-echo-postgres-hexa/internal/infra/api/handler"
	"github.com/andresxlp/go-echo-postgres-hexa/internal/infra/api/router"
	"github.com/andresxlp/go-echo-postgres-hexa/internal/infra/api/router/groups"
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

	_ = Container.Provide(postgresql.ConnectPGInstance)

	_ = Container.Provide(repo.NewUserRepository)

	_ = Container.Provide(handler.NewUserHandler)

	_ = Container.Provide(groups.NewUserRoutes)

	return Container
}
