package main

import (
	"silver-clean-code/internal/infra/db/local"
	"silver-clean-code/internal/infra/server"
	"silver-clean-code/internal/infra/server/echo"
	"silver-clean-code/internal/infra/server/gin"
)

const (
	EchoServer = "echo"
	GinServer  = "gin"
)

func main() {
	db := local.NewLocalDB()
	_ = db.CreateTable("accounts")

	manager := server.NewManager(db)

	http := factoryServer(EchoServer)
	http.Start(":8000", manager)
}

func factoryServer(name string) server.Server {
	switch name {
	case EchoServer:
		return echo.NewEchoServer()
	case GinServer:
		return gin.NewGinServer()
	}

	return nil
}
