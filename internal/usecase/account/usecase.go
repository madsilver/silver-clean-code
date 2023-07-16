package account

import "silver-clean-code/internal/entity"

type AccountUseCase struct {
	repository Repository
}

func NewAccountUseCase(repository Repository) *AccountUseCase {
	return &AccountUseCase{
		repository: repository,
	}
}

func (a *AccountUseCase) GetByID(id uint64) (*entity.Account, error) {
	return a.repository.FindByID(id)
}
