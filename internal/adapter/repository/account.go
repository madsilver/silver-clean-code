package repository

import (
	"silver-clean-code/internal/entity"
	"silver-clean-code/internal/infra/db/mysql"
)

type AccountRepository struct {
	db *mysql.MysqlDB
}

func NewAccountRepository(db *mysql.MysqlDB) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

func (r *AccountRepository) FindByID(id uint64) (*entity.Account, error) {
	account := &entity.Account{}
	row := r.db.Conn.QueryRow("SELECT * FROM Account WHERE AccountID = ?", id)
	err := row.Scan(&account.AccountID, &account.DocumentNumber)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (r *AccountRepository) FindAll() ([]*entity.Account, error) {
	account := &entity.Account{}
	var accounts []*entity.Account
	rows, err := r.db.Conn.Query("SELECT * FROM Account")
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		err = rows.Scan(&account.AccountID, &account.DocumentNumber)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (r *AccountRepository) Save(account *entity.Account) error {
	_, err := r.db.Conn.Exec("INSERT INTO Account (AccountID, DocumentNumber) VALUES (?, ?)", nil)
	return err
}
