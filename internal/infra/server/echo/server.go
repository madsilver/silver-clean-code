package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	e.Use(middleware.Logger())
	//e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	//	Format: "method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n",
	//}))

	Routes(e, manager)

	e.Logger.Fatal(e.Start(":" + env.GetString("SERVER_PORT", env.Port)))
}
