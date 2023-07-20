package server

import (
	"silver-clean-code/internal/adapter"
	"silver-clean-code/internal/adapter/controller"
	"silver-clean-code/internal/adapter/repository"
	"silver-clean-code/internal/infra/db"
	uca "silver-clean-code/internal/usecase/account"
)

type Controller interface {
	FindAccountByID(ctx adapter.ContextServer) error
	FindAccounts(ctx adapter.ContextServer) error
	CreateAccount(ctx adapter.ContextServer) error
}

type Manager struct {
	AccountController Controller
}

func NewManager(conn db.DB) *Manager {
	accountUseCase := uca.NewAccountUseCase(repository.NewAccountRepository(conn))

	return &Manager{
		AccountController: controller.NewAccountController(accountUseCase),
	}
}
