package controller

import (
	"errors"
	"github.com/golang/mock/gomock"
	"silver-clean-code/internal/adapter"
	mock_controller "silver-clean-code/internal/adapter/controller/mock"
	mock_adapter "silver-clean-code/internal/adapter/mock"
	"silver-clean-code/internal/entity"
	"testing"
)

func TestAccountController_FindByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctxMock := mock_adapter.NewMockContextServer(ctrl)
	ctxMock.EXPECT().Param(gomock.Any()).Return("1").Times(2)
	ctxMock.EXPECT().JSON(gomock.Any(), gomock.Any()).Return(nil).Times(2)
	usecaseMock := mock_controller.NewMockUseCase(ctrl)
	gomock.InOrder(
		usecaseMock.EXPECT().GetAccount(gomock.Any()).Return(&entity.Account{}, nil),
		usecaseMock.EXPECT().GetAccount(gomock.Any()).Return(nil, errors.New("error")),
	)
	type args struct {
		ctx adapter.ContextServer
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
	}
	c := NewAccountController(usecaseMock)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := c.FindAccountByID(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("FindByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccountController_FindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctxMock := mock_adapter.NewMockContextServer(ctrl)
	ctxMock.EXPECT().JSON(gomock.Any(), gomock.Any()).Return(nil).Times(2)
	usecaseMock := mock_controller.NewMockUseCase(ctrl)
	gomock.InOrder(
		usecaseMock.EXPECT().GetAccounts().Return([]*entity.Account{}, nil),
		usecaseMock.EXPECT().GetAccounts().Return(nil, errors.New("error")),
	)
	type args struct {
		ctx adapter.ContextServer
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
	c := NewAccountController(usecaseMock)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.FindAccounts(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("FindAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccountController_CreateAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctxMock := mock_adapter.NewMockContextServer(ctrl)
	usecaseMock := mock_controller.NewMockUseCase(ctrl)
	gomock.InOrder(
		//
		ctxMock.EXPECT().Bind(gomock.Any()).Return(errors.New("error")),
		ctxMock.EXPECT().JSON(gomock.Any(), gomock.Any()).Return(nil),
		//
		ctxMock.EXPECT().Bind(gomock.Any()).Return(nil),
		usecaseMock.EXPECT().SaveAccount(gomock.Any()).Return(errors.New("error")),
		ctxMock.EXPECT().JSON(gomock.Any(), gomock.Any()).Return(nil),
		//
		ctxMock.EXPECT().Bind(gomock.Any()).Return(nil),
		usecaseMock.EXPECT().SaveAccount(gomock.Any()).Return(nil),
		ctxMock.EXPECT().JSON(gomock.Any(), gomock.Any()).Return(nil),
	)
	type args struct {
		ctx adapter.ContextServer
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
			name:    "save account",
			args:    args{ctx: ctxMock},
			wantErr: false,
		},
	}
	c := NewAccountController(usecaseMock)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.CreateAccount(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
