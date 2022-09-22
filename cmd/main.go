package main

import (
	"fmt"

	"github.com/andresxlp/go-echo-postgres-hexa/cmd/providers"
	"github.com/andresxlp/go-echo-postgres-hexa/configs"
	"github.com/andresxlp/go-echo-postgres-hexa/internal/infra/api/router"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

var (
	serverHost = configs.Environments().ServerHost
	serverPort = configs.Environments().ServerPort
)

func main() {

	container := providers.BuildContainer()

	err := container.Invoke(func(router *router.Router, server *echo.Echo) {
		router.Init()
		server.Logger.Fatal(server.Start(fmt.Sprintf("%s:%d", serverHost, serverPort)))
	})

	if err != nil {
		log.Panic(err)
	}
}
