package server

import (
	"silver-clean-code/internal/adapter/controller/account"
	"silver-clean-code/internal/adapter/controller/transaction"
	uca "silver-clean-code/internal/usecase/account"
	uct "silver-clean-code/internal/usecase/transaction"
)

type Manager struct {
	AccountController     *account.AccountController
	TransactionController *transaction.TransactionController
}

func NewManager(accRepo uca.Repository, tranRepo uct.Repository) *Manager {
	accountUseCase := uca.NewAccountUseCase(accRepo)
	transactionUseCase := uct.NewTransactionUseCase(tranRepo)

	return &Manager{
		AccountController:     account.NewAccountController(accountUseCase),
		TransactionController: transaction.NewTransactionController(transactionUseCase),
	}
}

//func NewManager(conn db.DB) *Manager {
//	accountUseCase := uca.NewAccountUseCase(mysql.NewAccountRepository(conn))
//	transactionUseCase := uct.NewTransactionUseCase(mysql.NewTransactionRepository(conn))
//
//	return &Manager{
//		AccountController:     account.NewAccountController(accountUseCase),
//		TransactionController: transaction.NewTransactionController(transactionUseCase),
//	}
//}
