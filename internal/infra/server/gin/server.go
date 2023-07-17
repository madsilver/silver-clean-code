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
		_ = manager.AccountController.FindAccountByID(NewGinContext(c))
	})

	_ = g.Run(":" + port)
}
