package echo

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "silver-clean-code/docs"
	"silver-clean-code/internal/adapter/controller/core"
	"silver-clean-code/internal/infra/server"
)

// @title           Silver Clean Code
// @version         1.0
// @description     This project is a practical exercise in using the clean code pattern for application architecture.

// @contact.name   Rodrigo Prata
// @contact.email  rbpsilver@gmail.com

// @host      localhost:8000
// @BasePath  /v1

// Routes @externalDocs.description  OpenAPI
// @externalDocs.url          https://www.linkedin.com/in/silverdev/
func Routes(e *echo.Echo, manager *server.Manager) {
	e.GET("/v1/accounts/:id", mountHandler(manager.AccountController.FindAccountByID))
	e.GET("/v1/accounts", mountHandler(manager.AccountController.FindAccounts))
	e.POST("/v1/accounts", mountHandler(manager.AccountController.CreateAccount))
	e.GET("/v1/transactions/:id", mountHandler(manager.TransactionController.FindTransactionByID))
	e.GET("/v1/transactions", mountHandler(manager.TransactionController.FindTransactions))
	e.POST("/v1/transactions", mountHandler(manager.TransactionController.CreateTransaction))

	e.GET("/swagger/*", echoSwagger.WrapHandler)
}

func mountHandler(handler core.HandlerFunc) func(c echo.Context) error {
	return func(c echo.Context) error {
		return handler(NewContext(c))
	}
}
