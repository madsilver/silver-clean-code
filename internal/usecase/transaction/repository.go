package transaction

import "silver-clean-code/internal/entity"

type Repository interface {
	FindByID(id uint64) (*entity.Transaction, error)
	Save(transaction *entity.Transaction) error
}
