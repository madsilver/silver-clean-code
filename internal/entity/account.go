package entity

type Account struct {
	AccountID      uint64
	DocumentNumber string
}

func NewAccountEntity() *Account {
	return &Account{}
}
