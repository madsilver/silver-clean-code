package main

import (
	"silver-clean-code/internal/infra/db/local"
	"silver-clean-code/internal/infra/server"
)

func main() {
	db := local.NewLocalDB()
	_ = db.CreateTable("accounts")

	manager := server.NewManager(db)

	http := FactoryServer(ECHO)
	http.Start(":8000", manager)
}
