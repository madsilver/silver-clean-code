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
		return manager.AccountController.FindAccountByID(NewEchoContext(c))
	})

	e.GET("/accounts", func(c echo.Context) error {
		return manager.AccountController.FindAccounts(NewEchoContext(c))
	})

	e.POST("/accounts", func(c echo.Context) error {
		return manager.AccountController.CreateAccount(NewEchoContext(c))
	})

	e.Logger.Fatal(e.Start(":" + port))
}
