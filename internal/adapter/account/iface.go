package account

import (
	"silver-clean-code/internal/adapter"
	"silver-clean-code/internal/entity"
)

type UseCase interface {
	GetByID(id uint64) (*entity.Account, error)
}

type Controller interface {
	FindByID(ctx adapter.ContextServer) error
}
