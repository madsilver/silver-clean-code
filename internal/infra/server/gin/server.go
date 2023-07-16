package gin

import (
	"github.com/gin-gonic/gin"
	"silver-clean-code/internal/infra/server"
)

type GinServer struct {
}

func NewGinServer() *GinServer {
	return &GinServer{}
}

func (s *GinServer) Start(port string, manager *server.Manager) {
	g := gin.Default()

	g.GET("/accounts/:id", func(c *gin.Context) {
		_ = manager.AccountController.FindByID(NewGinContext(c))
	})

	//g.POST("/accounts", func(c *gin.Context) {
	//	_ = manager.AccountController.Create(NewGinContext(c))
	//})

	_ = g.Run(":" + port)
}
