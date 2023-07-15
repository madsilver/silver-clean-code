package account

import "silver-clean-code/internal/entity"

type Account struct {
	AccountID      uint64 `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

func EntityToPresenter(ent *entity.Account) Account {
	return Account{
		AccountID:      ent.AccountID,
		DocumentNumber: ent.DocumentNumber,
	}
}
