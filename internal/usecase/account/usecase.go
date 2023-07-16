package account

import "silver-clean-code/internal/entity"

type Repository interface {
	FindByID(id uint64) (*entity.Account, error)
	FindAll() ([]*entity.Account, error)
}

type AccountUseCase struct {
	repository Repository
}

func NewAccountUseCase(repository Repository) *AccountUseCase {
	return &AccountUseCase{
		repository: repository,
	}
}

func (a *AccountUseCase) GetAccount(id uint64) (*entity.Account, error) {
	return a.repository.FindByID(id)
}

func (a *AccountUseCase) GetAccounts() ([]*entity.Account, error) {
	return a.repository.FindAll()
}
