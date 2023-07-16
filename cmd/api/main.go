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

	serverName := env.GetString("SERVER_NAME", env.EchoServer)
	http := factoryServer(serverName)
	http.Start(env.GetString("SERVER_PORT", env.Port), manager)
}

func factoryServer(name string) server.Server {
	switch name {
	case env.EchoServer:
		return echo.NewEchoServer()
	case env.GinServer:
		return gin.NewGinServer()
	}

	return nil
}
