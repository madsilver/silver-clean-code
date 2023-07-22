package gin

import (
	"github.com/gin-gonic/gin"
	"silver-clean-code/internal/adapter"
	"silver-clean-code/internal/infra/server"
)

func Routes(g *gin.Engine, manager *server.Manager) {
	g.GET("/accounts/:id", mountHandler(manager.AccountController.FindAccountByID))
	g.GET("/accounts", mountHandler(manager.AccountController.FindAccounts))
	g.POST("/accounts", mountHandler(manager.AccountController.CreateAccount))
	g.GET("/transactions", mountHandler(manager.TransactionController.FindTransactionByID))
	g.POST("/transactions", mountHandler(manager.TransactionController.CreateTransaction))
}

func mountHandler(fnController func(ctx adapter.ContextServer) error) func(c *gin.Context) {
	return func(g *gin.Context) {
		ct := NewContext(g)
		_ = fnController(ct)
	}
}
