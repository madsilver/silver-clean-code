package transaction

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"silver-clean-code/internal/entity"
	mock_transaction "silver-clean-code/internal/usecase/transaction/mock"
	"testing"
)

func TestTransactionUseCase_SaveTransaction(t *testing.T) {
	tests := []struct {
		name        string
		transaction *entity.Transaction
		wantErr     bool
		wantAmount  float64
	}{
		{
			name: "should save buy cache transaction",
			transaction: &entity.Transaction{
				AccountID:       1,
				Amount:          10,
				OperationTypeID: entity.OperationType.BuyCash,
			},
			wantErr:    false,
			wantAmount: -10.0,
		},
		{
			name: "should save payment transaction",
			transaction: &entity.Transaction{
				AccountID:       1,
				Amount:          -10,
				OperationTypeID: entity.OperationType.Payment,
			},
			wantErr:    false,
			wantAmount: 10.0,
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mock_transaction.NewMockRepository(ctrl)
	repository.EXPECT().Save(gomock.Any()).AnyTimes()

	uc := NewTransactionUseCase(repository)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := uc.SaveTransaction(tt.transaction)
			assert.Nil(t, err)
			assert.NotNil(t, tt.transaction.EventDate)
			assert.Equal(t, tt.wantAmount, tt.transaction.Amount)
		})
	}
}

func TestTransactionUseCase_GetTransaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mock_transaction.NewMockRepository(ctrl)
	repository.EXPECT().FindByID(gomock.Any()).Return(&entity.Transaction{}, nil)
	uc := NewTransactionUseCase(repository)

	transaction, err := uc.GetTransaction(1)

	assert.Nil(t, err)
	assert.NotNil(t, transaction)

}
