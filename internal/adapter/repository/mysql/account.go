package mysql

import (
	"silver-clean-code/internal/entity"
	"silver-clean-code/internal/infra/db"
)

type AccountRepository struct {
	db db.DB
}

func NewAccountRepository(db db.DB) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

func (ar *AccountRepository) FindByID(id uint64) (*entity.Account, error) {
	query := "SELECT * FROM Account WHERE AccountID = ?"
	account := &entity.Account{}
	err := ar.db.QueryRow(query, id, func(scan func(dest ...any) error) error {
		return scan(&account.AccountID, &account.DocumentNumber)
	})
	return account, err
}

func (ar *AccountRepository) FindAll() ([]*entity.Account, error) {
	query := "SELECT * FROM Account"
	var accounts []*entity.Account

	err := ar.db.Query(query, func(scan func(dest ...any) error) error {
		account := &entity.Account{}
		err := scan(&account.AccountID, &account.DocumentNumber)
		if err == nil {
			accounts = append(accounts, account)
		}
		return err
	})
	return accounts, err
}

func (ar *AccountRepository) Save(account *entity.Account) error {
	query := "INSERT INTO Account (DocumentNumber) VALUES (?)"
	res, err := ar.db.Save(query, &account.DocumentNumber)
	if err == nil {
		account.AccountID = uint64(res.(int64))
	}
	return err
}
