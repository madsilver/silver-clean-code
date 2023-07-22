package main

import (
	"silver-clean-code/internal/infra/db/mysql"
	"silver-clean-code/internal/infra/env"
	"silver-clean-code/internal/infra/server"
	"silver-clean-code/internal/infra/server/echo"
	"silver-clean-code/internal/infra/server/gin"
)

func main() {
	db := mysql.NewMysqlDB()

	manager := server.NewManager(db)

	factoryServer().Start(manager)
}

func factoryServer() server.Server {
	serverName := env.GetString("SERVER_NAME", env.EchoServer)
	switch serverName {
	case env.EchoServer:
		return echo.NewEchoServer()
	case env.GinServer:
		return gin.NewGinServer()
	}
	return nil
}
