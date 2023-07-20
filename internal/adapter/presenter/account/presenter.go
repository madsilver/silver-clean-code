package account

import "silver-clean-code/internal/entity"

type Account struct {
	AccountID      uint64 `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

type Accounts struct {
	TotalItems int       `json:"total_items"`
	Items      []Account `json:"items"`
}

func ToEntity(pre *Account) *entity.Account {
	return &entity.Account{
		AccountID:      pre.AccountID,
		DocumentNumber: pre.DocumentNumber,
	}
}

func ToPresenter(ent *entity.Account) Account {
	return Account{
		AccountID:      ent.AccountID,
		DocumentNumber: ent.DocumentNumber,
	}
}

func ToPresenters(entities []*entity.Account) Accounts {
	var presenters []Account
	for _, ent := range entities {
		presenter := Account{
			AccountID:      ent.AccountID,
			DocumentNumber: ent.DocumentNumber,
		}
		presenters = append(presenters, presenter)
	}
	return Accounts{
		Items:      presenters,
		TotalItems: len(presenters),
	}
}
