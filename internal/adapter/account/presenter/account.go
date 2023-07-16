package presenter

import "silver-clean-code/internal/entity"

type Account struct {
	AccountID      uint64 `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

type Accounts struct {
	TotalItems int       `json:"total_items"`
	Items      []Account `json:"items"`
}

func EntityToPresenter(ent *entity.Account) Account {
	return Account{
		AccountID:      ent.AccountID,
		DocumentNumber: ent.DocumentNumber,
	}
}

func EntitiesToPresenters(entities []*entity.Account) Accounts {
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
