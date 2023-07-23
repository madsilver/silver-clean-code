package presenter

import "silver-clean-code/internal/entity"

type Account struct {
	AccountID      uint64 `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

func NewAccountPresenter(ent *entity.Account) *Account {
	account := &Account{}
	if ent != nil {
		return account.Parse(ent)
	}
	return account
}

func (ac *Account) ToEntity() *entity.Account {
	return &entity.Account{
		AccountID:      ac.AccountID,
		DocumentNumber: ac.DocumentNumber,
	}
}

func (ac *Account) Parse(ent *entity.Account) *Account {
	ac.AccountID = ent.AccountID
	ac.DocumentNumber = ent.DocumentNumber
	return ac
}
