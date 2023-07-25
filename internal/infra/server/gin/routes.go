package gin

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "silver-clean-code/docs"
	"silver-clean-code/internal/adapter/controller"
	"silver-clean-code/internal/infra/server"
)

func Routes(g *gin.Engine, manager *server.Manager) {
	g.GET("/v1/accounts/:id", mountHandler(manager.AccountController.FindAccountByID))
	g.GET("/v1/accounts", mountHandler(manager.AccountController.FindAccounts))
	g.POST("/v1/accounts", mountHandler(manager.AccountController.CreateAccount))
	g.GET("/v1/transactions/:id", mountHandler(manager.TransactionController.FindTransactionByID))
	g.GET("/v1/transactions", mountHandler(manager.TransactionController.FindTransactions))
	g.POST("/v1/transactions", mountHandler(manager.TransactionController.CreateTransaction))

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func mountHandler(fnController func(ctx controller.ContextServer) error) func(c *gin.Context) {
	return func(g *gin.Context) {
		ct := NewContext(g)
		_ = fnController(ct)
	}
}
