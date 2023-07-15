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
	r := gin.Default()

	r.GET("/accounts/:id", func(c *gin.Context) {
		_ = manager.AccountController.FindByID(NewGinContext(c))
	})

	_ = r.Run(port)
}
