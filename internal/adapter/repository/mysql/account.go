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

func (r *AccountRepository) FindByID(id uint64) (*entity.Account, error) {
	query := "SELECT * FROM Account WHERE AccountID = ?"
	account := &entity.Account{}
	err := r.db.QueryRow(query, id, func(scan func(dest ...any) error) {
		_ = scan(&account.AccountID, &account.DocumentNumber)
	})
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (r *AccountRepository) FindAll() ([]*entity.Account, error) {
	query := "SELECT * FROM Account"
	var accounts []*entity.Account

	err := r.db.Query(query, func(scan func(dest ...any) error) {
		account := &entity.Account{}
		err := scan(&account.AccountID, &account.DocumentNumber)
		if err == nil {
			accounts = append(accounts, account)
		}
	})

	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (r *AccountRepository) Save(account *entity.Account) error {
	query := "INSERT INTO Account (DocumentNumber) VALUES (?)"
	res, err := r.db.Save(query, &account.DocumentNumber)
	if err != nil {
		return err
	}
	account.AccountID = uint64(res.(int64))
	return err
}
