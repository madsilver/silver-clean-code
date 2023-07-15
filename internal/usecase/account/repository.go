package account

import "silver-clean-code/internal/entity"

type AccountRepository interface {
	FindByID(id uint64) (*entity.Account, error)
}
