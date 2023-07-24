package main

import (
	"github.com/labstack/gommon/log"
	mysqlAdapter "silver-clean-code/internal/adapter/repository/mysql"
	postgresAdapter "silver-clean-code/internal/adapter/repository/postgres"
	"silver-clean-code/internal/infra/db"
	mysqlDB "silver-clean-code/internal/infra/db/mysql"
	postgresDB "silver-clean-code/internal/infra/db/postgres"
	"silver-clean-code/internal/infra/env"
	"silver-clean-code/internal/infra/server"
	"silver-clean-code/internal/infra/server/echo"
	"silver-clean-code/internal/infra/server/gin"
	"silver-clean-code/internal/usecase/account"
	"silver-clean-code/internal/usecase/transaction"
)

func main() {
	dbApp := FactoryDB()

	manager := server.NewManager(FactoryRepositories(dbApp))

	FactoryServer().Start(manager)
}

func FactoryServer() server.Server {
	serverName := env.GetString("SERVER_NAME", env.EchoServer)

	switch serverName {
	case env.EchoServer:
		log.Info("Http server: Echo")
		return echo.NewEchoServer()
	case env.GinServer:
		log.Info("Http server: Gin")
		return gin.NewGinServer()
	}
	panic("undefined server")
}

func FactoryDB() db.DB {
	dbServerName := env.GetString("DB_NAME", env.MysqlServer)

	switch dbServerName {
	case env.MysqlServer:
		log.Info("Database: MySQL")
		return mysqlDB.NewMysqlDB()
	case env.PostgresServer:
		log.Info("Database: Postgres")
		return postgresDB.NewPostgresDB()
	}
	panic("undefined database")
}

func FactoryRepositories(db db.DB) (account.Repository, transaction.Repository) {
	dbServerName := env.GetString("DB_NAME", env.MysqlServer)

	switch dbServerName {
	case env.MysqlServer:
		return mysqlAdapter.NewAccountRepository(db), mysqlAdapter.NewTransactionRepository(db)
	case env.PostgresServer:
		return postgresAdapter.NewAccountRepository(db), nil
	}
	panic("undefined repository")
}
