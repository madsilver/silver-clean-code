package mysql

import (
	"silver-clean-code/internal/entity"
	"silver-clean-code/internal/infra/db"
)

type TransactionRepository struct {
	db db.DB
}

func NewTransactionRepository(db db.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (r *TransactionRepository) FindByID(id uint64) (*entity.Transaction, error) {
	query := "SELECT * FROM Transaction WHERE TransactionID = ?"
	transaction := &entity.Transaction{}
	err := r.db.QueryRow(query, id, func(scan func(dest ...any) error) {
		_ = scan(
			&transaction.TransactionID,
			&transaction.AccountID,
			&transaction.OperationTypeID,
			&transaction.Amount,
			&transaction.EventDate,
		)
	})
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r *TransactionRepository) Save(transaction *entity.Transaction) error {
	query := "INSERT INTO Transaction (AccountID, OperationTypeID, Amount, EventDate) VALUES (?,?,?,?)"
	res, err := r.db.Save(query, &transaction.AccountID, &transaction.OperationTypeID, &transaction.Amount, &transaction.EventDate)
	if err != nil {
		return err
	}
	transaction.TransactionID = uint64(res.(int64))
	return err
}
