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
			name:    "",
			args:    args{ctx: ctxMock},
			wantErr: false,
		},
		{
			name:    "",
			args:    args{ctx: ctxMock},
			wantErr: false,
		},
	}
	c := NewAccountController(usecaseMock)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := c.FindByID(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("FindByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
