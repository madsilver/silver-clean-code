package server

import (
	"silver-clean-code/internal/adapter"
	"silver-clean-code/internal/adapter/account"
	uca "silver-clean-code/internal/usecase/account"
)

type Manager struct {
	AccountController Controller
}

func NewManager(db adapter.DB) *Manager {
	accountUseCase := uca.NewAccountUseCase(account.NewAccountRepository(db))

	return &Manager{
		AccountController: account.NewAccountController(accountUseCase),
	}
}
