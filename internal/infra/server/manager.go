package server

import (
	"silver-clean-code/internal/adapter"
	"silver-clean-code/internal/adapter/controller"
	"silver-clean-code/internal/adapter/repository"
	"silver-clean-code/internal/infra/db/mysql"
	uca "silver-clean-code/internal/usecase/account"
)

type Controller interface {
	FindByID(ctx adapter.ContextServer) error
	FindAll(ctx adapter.ContextServer) error
}

type Manager struct {
	AccountController Controller
}

func NewManager(conn *mysql.MysqlDB) *Manager {
	accountUseCase := uca.NewAccountUseCase(repository.NewAccountRepository(conn))

	return &Manager{
		AccountController: controller.NewAccountController(accountUseCase),
	}
}
