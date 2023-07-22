package echo

import (
	"github.com/labstack/echo/v4"
	"silver-clean-code/internal/infra/env"
	"silver-clean-code/internal/infra/server"
)

type EchoServer struct {
}

func NewEchoServer() *EchoServer {
	return &EchoServer{}
}

func (s *EchoServer) Start(manager *server.Manager) {
	e := echo.New()

	Routes(e, manager)

	e.Logger.Fatal(e.Start(":" + env.GetString("SERVER_PORT", env.Port)))
}
