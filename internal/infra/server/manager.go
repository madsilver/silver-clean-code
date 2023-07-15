package server

import (
	"silver-clean-code/internal/adapter"
	"silver-clean-code/internal/adapter/account"
	uca "silver-clean-code/internal/usecase/account"
)

type Manager struct {
	AccountController *account.AccountController
}

func NewManager(repository adapter.Repository) *Manager {
	return &Manager{
		AccountController: account.NewAccountController(accountUseCase(repository)),
	}
}

func accountUseCase(repository adapter.Repository) *uca.AccountUseCase {
	accountRepository := account.NewAccountRepository(repository)
	return uca.NewAccountUseCase(accountRepository)
}
