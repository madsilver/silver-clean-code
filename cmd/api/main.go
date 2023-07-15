package main

import (
	"silver-clean-code/internal/infra/database"
	"silver-clean-code/internal/infra/server"
	"silver-clean-code/internal/infra/server/echo"
)

func main() {
	localDB := database.NewLocalDB()
	_ = localDB.CreateTable("accounts")

	manager := server.NewManager(localDB)

	echo.Start(manager, ":8000")
	//gin.Start(manager, "8001")
}
