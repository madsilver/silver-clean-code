package main

import (
	"silver-clean-code/internal/infra/server"
)

func main() {
	dbApp := factoryDB()

	manager := server.NewManager(factoryRepositories(dbApp))

	factoryServer().Start(manager)
}
