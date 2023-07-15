package echo

import (
	"github.com/labstack/echo/v4"
	"silver-clean-code/internal/infra/server"
)

func Start(manager *server.Manager, port string) {
	e := echo.New()

	e.GET("/accounts/:id", func(c echo.Context) error {
		return manager.AccountController.FindByID(NewEchoContext(c))
	})

	e.Logger.Fatal(e.Start(port))
}
