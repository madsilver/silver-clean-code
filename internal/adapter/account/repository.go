package account

import (
	"silver-clean-code/internal/adapter"
	"silver-clean-code/internal/entity"
)

type AccountRepository struct {
	db        adapter.DB
	tableName string
}

func NewAccountRepository(db adapter.DB) *AccountRepository {
	return &AccountRepository{
		db:        db,
		tableName: "accounts",
	}
}

func (r *AccountRepository) FindByID(id uint64) (*entity.Account, error) {
	account, err := r.db.FindByID(r.tableName, id)
	if err != nil {
		return nil, err
	}
	return account.(*entity.Account), nil
}
