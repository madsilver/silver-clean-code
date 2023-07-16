package echo

import (
	"github.com/labstack/echo/v4"
	"silver-clean-code/internal/infra/server"
)

type EchoServer struct {
}

func NewEchoServer() *EchoServer {
	return &EchoServer{}
}

func (s *EchoServer) Start(port string, manager *server.Manager) {
	e := echo.New()

	e.GET("/accounts/:id", func(c echo.Context) error {
		return manager.AccountController.FindByID(NewEchoContext(c))
	})

	e.GET("/accounts", func(c echo.Context) error {
		return manager.AccountController.FindAll(NewEchoContext(c))
	})

	e.Logger.Fatal(e.Start(":" + port))
}
