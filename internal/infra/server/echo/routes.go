package echo

import (
	"github.com/labstack/echo/v4"
	"silver-clean-code/internal/adapter"
	"silver-clean-code/internal/infra/server"
)

func Routes(e *echo.Echo, manager *server.Manager) {
	e.GET("/accounts/:id", mountHandler(manager.AccountController.FindAccountByID))
	e.GET("/accounts", mountHandler(manager.AccountController.FindAccounts))
	e.POST("/accounts", mountHandler(manager.AccountController.CreateAccount))
	e.GET("/transactions/:id", mountHandler(manager.TransactionController.FindTransactionByID))
	e.GET("/transactions", mountHandler(manager.TransactionController.FindTransactions))
	e.POST("/transactions", mountHandler(manager.TransactionController.CreateTransaction))
}

func mountHandler(fnController func(ctx adapter.ContextServer) error) func(c echo.Context) error {
	return func(c echo.Context) error {
		ct := NewContext(c)
		return fnController(ct)
	}
}
