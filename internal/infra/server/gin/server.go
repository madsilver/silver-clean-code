package gin

import (
	"github.com/gin-gonic/gin"
	"silver-clean-code/internal/infra/server"
)

func Start(manager *server.Manager, port string) {
	r := gin.Default()

	r.GET("/accounts/:id", func(c *gin.Context) {
		_ = manager.AccountController.FindByID(NewGinContext(c))
	})

	_ = r.Run(port)
}
