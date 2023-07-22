package gin

import (
	"github.com/gin-gonic/gin"
	"silver-clean-code/internal/infra/env"
	"silver-clean-code/internal/infra/server"
)

type GinServer struct {
}

func NewGinServer() *GinServer {
	return &GinServer{}
}

func (s *GinServer) Start(manager *server.Manager) {
	g := gin.Default()

	Routes(g, manager)

	_ = g.Run(":" + env.GetString("SERVER_PORT", env.Port))
}
