package postgres

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

func (tr *AccountRepository) FindByID(id uint64) (*entity.Account, error) {
	query := "SELECT * FROM \"Account\" WHERE \"AccountID\" = $1"
	account := &entity.Account{}
	err := tr.db.QueryRow(query, id, func(scan func(dest ...any) error) error {
		return scan(&account.AccountID, &account.DocumentNumber)
	})
	return account, err
}

func (tr *AccountRepository) FindAll() ([]*entity.Account, error) {
	query := "SELECT * FROM \"Account\""
	var accounts []*entity.Account

	err := tr.db.Query(query, func(scan func(dest ...any) error) error {
		account := &entity.Account{}
		err := scan(&account.AccountID, &account.DocumentNumber)
		if err == nil {
			accounts = append(accounts, account)
		}
		return err
	})
	return accounts, err
}

func (tr *AccountRepository) Save(account *entity.Account) error {
	query := "INSERT INTO \"Account\" (\"DocumentNumber\") VALUES ($1) RETURNING \"AccountID\""
	res, err := tr.db.Save(query, &account.DocumentNumber)
	if err == nil {
		account.AccountID = uint64(res.(int64))
	}
	return err
}
