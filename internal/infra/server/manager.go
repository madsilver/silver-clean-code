package server

import (
	"silver-clean-code/internal/adapter"
	"silver-clean-code/internal/adapter/account"
	uca "silver-clean-code/internal/usecase/account"
)

type Manager struct {
	AccountController *account.AccountController
}

func NewManager(db adapter.DB) *Manager {
	return &Manager{
		AccountController: account.NewAccountController(accountUseCase(db)),
	}
}

func accountUseCase(db adapter.DB) *uca.AccountUseCase {
	accountRepository := account.NewAccountRepository(db)
	return uca.NewAccountUseCase(accountRepository)
}
