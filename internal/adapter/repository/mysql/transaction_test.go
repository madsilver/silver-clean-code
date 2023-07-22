package mysql

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"silver-clean-code/internal/entity"
	mock_db "silver-clean-code/internal/infra/db/mock"
	"testing"
	"time"
)

func TestTransactionRepository_FindByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	data := &entity.Transaction{
		TransactionID:   1,
		AccountID:       1,
		OperationTypeID: entity.OperationType.Payment,
		Amount:          100.0,
		EventDate:       time.Now().Format(time.RFC3339Nano),
	}
	qfn := buildTransactionQueryRowFunction(data)
	mockDB := mock_db.NewMockDB(ctrl)
	mockDB.EXPECT().
		QueryRow(gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(qfn)
	repository := NewTransactionRepository(mockDB)

	transaction, err := repository.FindByID(uint64(1))

	assert.Nil(t, err)
	assert.Equal(t, uint64(1), transaction.TransactionID)
}

func TestTransactionRepository_FindByID_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mock_db.NewMockDB(ctrl)
	mockDB.EXPECT().
		QueryRow(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(errors.New("error"))
	repository := NewTransactionRepository(mockDB)

	transaction, err := repository.FindByID(uint64(1))

	assert.NotNil(t, err)
	assert.Nil(t, transaction)
}

func TestTransactionRepository_Save(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mock_db.NewMockDB(ctrl)
	mockDB.EXPECT().
		Save(gomock.Any(), gomock.Any()).
		Return(int64(1), nil)
	repository := NewTransactionRepository(mockDB)

	err := repository.Save(&entity.Transaction{})

	assert.Nil(t, err)
}

func TestTransactionRepository_Save_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mock_db.NewMockDB(ctrl)
	mockDB.EXPECT().
		Save(gomock.Any(), gomock.Any()).
		Return(nil, errors.New("error"))
	repository := NewTransactionRepository(mockDB)

	err := repository.Save(&entity.Transaction{})

	assert.NotNil(t, err)
}

func buildTransactionQueryFunction(transaction *entity.Transaction) any {
	return func(qs string, query func(scan func(dest ...any) error)) error {
		query(func(dest ...any) error {
			*dest[0].(*uint64) = transaction.TransactionID
			*dest[1].(*uint64) = transaction.AccountID
			*dest[2].(*int) = transaction.OperationTypeID
			*dest[3].(*float64) = transaction.Amount
			*dest[4].(*string) = transaction.EventDate

			return nil
		})
		return nil
	}
}

func buildTransactionQueryRowFunction(transaction *entity.Transaction) any {
	return func(qs string, id uint64, query func(scan func(dest ...any) error)) error {
		query(func(dest ...any) error {
			*dest[0].(*uint64) = transaction.TransactionID
			*dest[1].(*uint64) = transaction.AccountID
			*dest[2].(*int) = transaction.OperationTypeID
			*dest[3].(*float64) = transaction.Amount
			*dest[4].(*string) = transaction.EventDate
			return nil
		})
		return nil
	}
}
