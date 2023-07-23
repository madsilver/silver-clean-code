package transaction

import (
	"silver-clean-code/internal/entity"
)

type TransactionUseCase struct {
	repository Repository
}

func NewTransactionUseCase(repository Repository) *TransactionUseCase {
	return &TransactionUseCase{
		repository: repository,
	}
}

func (a *TransactionUseCase) GetTransaction(id uint64) (*entity.Transaction, error) {
	return a.repository.FindByID(id)
}

func (a *TransactionUseCase) GetTransactions() ([]*entity.Transaction, error) {
	return a.repository.FindAll()
}

func (a *TransactionUseCase) SaveTransaction(transaction *entity.Transaction) error {
	transaction.TimeNow().NormalizeAmount()
	return a.repository.Save(transaction)
}
