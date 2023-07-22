package main

import (
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

func factoryServer() server.Server {
	serverName := env.GetString("SERVER_NAME", env.EchoServer)

	switch serverName {
	case env.EchoServer:
		return echo.NewEchoServer()
	case env.GinServer:
		return gin.NewGinServer()
	}
	panic("undefined server")
}

func factoryDB() db.DB {
	dbServerName := env.GetString("DB_SERVER", env.MysqlServer)

	switch dbServerName {
	case env.MysqlServer:
		return mysqlDB.NewMysqlDB()
	case env.PostgresServer:
		return postgresDB.NewPostgresDB()
	}
	panic("undefined database")
}

func factoryRepositories(db db.DB) (account.Repository, transaction.Repository) {
	dbServerName := env.GetString("DB_SERVER", env.MysqlServer)

	switch dbServerName {
	case env.MysqlServer:
		return mysqlAdapter.NewAccountRepository(db), mysqlAdapter.NewTransactionRepository(db)
	case env.PostgresServer:
		return postgresAdapter.NewAccountRepository(db), nil
	}
	panic("undefined repository")
}
