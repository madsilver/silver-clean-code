package repository

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"silver-clean-code/internal/entity"
	mock_db "silver-clean-code/internal/infra/db/mock"
	"testing"
)

func TestAccountRepository_FindByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id := uint64(1)
	doc := "00000000001"
	qfn := buildQueryRowFunction(id, doc, false)
	mockDB := mock_db.NewMockDB(ctrl)
	mockDB.EXPECT().
		QueryRow(gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(qfn)
	repository := NewAccountRepository(mockDB)

	account, err := repository.FindByID(uint64(1))

	assert.Nil(t, err)
	assert.Equal(t, id, account.AccountID)
	assert.Equal(t, doc, account.DocumentNumber)
}

func TestAccountRepository_FindByID_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	qfn := buildQueryRowFunction(0, "", true)
	mockDB := mock_db.NewMockDB(ctrl)
	mockDB.EXPECT().
		QueryRow(gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(qfn)
	repository := NewAccountRepository(mockDB)

	account, err := repository.FindByID(uint64(1))

	assert.NotNil(t, err)
	assert.Nil(t, account)
}

func TestAccountRepository_FindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id := uint64(1)
	doc := "00000000001"
	qfn := buildQueryFunction(id, doc, false)
	mockDB := mock_db.NewMockDB(ctrl)
	mockDB.EXPECT().
		Query(gomock.Any(), gomock.Any()).
		DoAndReturn(qfn)
	repository := NewAccountRepository(mockDB)

	accounts, err := repository.FindAll()

	assert.Nil(t, err)
	assert.Equal(t, 1, len(accounts))
	assert.Equal(t, id, accounts[0].AccountID)
	assert.Equal(t, doc, accounts[0].DocumentNumber)
}

func TestAccountRepository_FindAll_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	qfn := buildQueryFunction(0, "", true)
	mockDB := mock_db.NewMockDB(ctrl)
	mockDB.EXPECT().
		Query(gomock.Any(), gomock.Any()).
		DoAndReturn(qfn)
	repository := NewAccountRepository(mockDB)

	accounts, err := repository.FindAll()

	assert.NotNil(t, err)
	assert.Nil(t, accounts)
}

func TestAccountRepository_Save(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mock_db.NewMockDB(ctrl)
	mockDB.EXPECT().
		Save(gomock.Any(), gomock.Any()).
		Return(int64(1), nil)
	repository := NewAccountRepository(mockDB)

	err := repository.Save(&entity.Account{})

	assert.Nil(t, err)
}

func TestAccountRepository_Save_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mock_db.NewMockDB(ctrl)
	mockDB.EXPECT().
		Save(gomock.Any(), gomock.Any()).
		Return(int64(0), errors.New("error"))
	repository := NewAccountRepository(mockDB)

	err := repository.Save(&entity.Account{})

	assert.NotNil(t, err)
}

func buildQueryFunction(id uint64, doc string, wantErr bool) any {
	return func(qs string, query func(scan func(dest ...any) error)) error {
		query(func(dest ...any) error {
			*dest[0].(*uint64) = id
			*dest[1].(*string) = doc
			return nil
		})
		if wantErr {
			return errors.New("error")
		}
		return nil
	}
}

func buildQueryRowFunction(id uint64, doc string, wantErr bool) any {
	return func(qs string, id uint64, query func(scan func(dest ...any) error)) error {
		query(func(dest ...any) error {
			*dest[0].(*uint64) = id
			*dest[1].(*string) = doc
			return nil
		})
		if wantErr {
			return errors.New("error")
		}
		return nil
	}
}
