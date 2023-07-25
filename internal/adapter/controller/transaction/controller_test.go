package transaction

import (
	"errors"
	"github.com/golang/mock/gomock"
	"silver-clean-code/internal/adapter/controller"
	mock_controller "silver-clean-code/internal/adapter/controller/mock"
	mock_transaction "silver-clean-code/internal/adapter/controller/transaction/mock"
	"silver-clean-code/internal/entity"
	"testing"
)

func TestTransactionController_CreateTransaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctxMock := mock_controller.NewMockContextServer(ctrl)
	usecaseMock := mock_transaction.NewMockUseCase(ctrl)
	gomock.InOrder(
		//
		ctxMock.EXPECT().Bind(gomock.Any()).Return(errors.New("error")),
		ctxMock.EXPECT().JSON(gomock.Any(), gomock.Any()).Return(nil),
		//
		ctxMock.EXPECT().Bind(gomock.Any()).Return(nil),
		usecaseMock.EXPECT().SaveTransaction(gomock.Any()).Return(errors.New("error")),
		ctxMock.EXPECT().JSON(gomock.Any(), gomock.Any()).Return(nil),
		//
		ctxMock.EXPECT().Bind(gomock.Any()).Return(nil),
		usecaseMock.EXPECT().SaveTransaction(gomock.Any()).Return(nil),
		ctxMock.EXPECT().JSON(gomock.Any(), gomock.Any()).Return(nil),
	)
	type args struct {
		ctx controller.ContextServer
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "bad request",
			args:    args{ctx: ctxMock},
			wantErr: false,
		},
		{
			name:    "internal server error",
			args:    args{ctx: ctxMock},
			wantErr: false,
		},
		{
			name:    "save Transaction",
			args:    args{ctx: ctxMock},
			wantErr: false,
		},
	}
	c := NewTransactionController(usecaseMock)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.CreateTransaction(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("CreateTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTransactionController_FindTransactionByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctxMock := mock_controller.NewMockContextServer(ctrl)
	ctxMock.EXPECT().Param(gomock.Any()).Return("1").Times(2).AnyTimes()
	ctxMock.EXPECT().JSON(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	usecaseMock := mock_transaction.NewMockUseCase(ctrl)
	gomock.InOrder(
		usecaseMock.EXPECT().GetTransaction(gomock.Any()).Return(&entity.Transaction{TransactionID: 1}, nil),
		usecaseMock.EXPECT().GetTransaction(gomock.Any()).Return(&entity.Transaction{TransactionID: 0}, nil),
		usecaseMock.EXPECT().GetTransaction(gomock.Any()).Return(nil, errors.New("error")),
	)
	type args struct {
		ctx controller.ContextServer
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "should return status ok",
			args:    args{ctx: ctxMock},
			wantErr: false,
		},
		{
			name:    "should return status not found",
			args:    args{ctx: ctxMock},
			wantErr: false,
		},
		{
			name:    "should return internal server error",
			args:    args{ctx: ctxMock},
			wantErr: false,
		},
	}
	c := NewTransactionController(usecaseMock)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := c.FindTransactionByID(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("FindByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTransactionController_FindTransactions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctxMock := mock_controller.NewMockContextServer(ctrl)
	ctxMock.EXPECT().JSON(gomock.Any(), gomock.Any()).Return(nil).Times(2)
	usecaseMock := mock_transaction.NewMockUseCase(ctrl)
	gomock.InOrder(
		usecaseMock.EXPECT().GetTransactions().Return([]*entity.Transaction{}, nil),
		usecaseMock.EXPECT().GetTransactions().Return(nil, errors.New("error")),
	)
	type args struct {
		ctx controller.ContextServer
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "should return status ok",
			args:    args{ctx: ctxMock},
			wantErr: false,
		},
		{
			name:    "should return internal server error",
			args:    args{ctx: ctxMock},
			wantErr: false,
		},
	}
	c := NewTransactionController(usecaseMock)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.FindTransactions(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("FindTransactions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
