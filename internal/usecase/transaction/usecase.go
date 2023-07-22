package transaction

import (
	"math"
	"silver-clean-code/internal/entity"
	"time"
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
	setAmountValueByType(transaction)
	setEventDate(transaction)
	return a.repository.Save(transaction)
}

func setAmountValueByType(transaction *entity.Transaction) {
	transaction.Amount = math.Abs(transaction.Amount)
	if transaction.OperationTypeID != entity.OperationType.Payment {
		transaction.Amount = transaction.Amount * -1
	}
}

func setEventDate(transaction *entity.Transaction) {
	t := time.Now()
	transaction.EventDate = t.Format(time.RFC3339Nano)
}
