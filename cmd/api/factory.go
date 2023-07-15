package main

import (
	"silver-clean-code/internal/infra/server"
	"silver-clean-code/internal/infra/server/echo"
	"silver-clean-code/internal/infra/server/gin"
)

const (
	ECHO = "echo"
	GIN  = "gin"
)

func FactoryServer(name string) server.Server {
	switch name {
	case ECHO:
		return echo.NewEchoServer()
	case GIN:
		return gin.NewGinServer()
	}

	return nil
}
