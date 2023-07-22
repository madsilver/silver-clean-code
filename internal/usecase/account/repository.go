package account

import "silver-clean-code/internal/entity"

type Repository interface {
	FindByID(id uint64) (*entity.Account, error)
	FindAll() ([]*entity.Account, error)
	Save(account *entity.Account) error
}
