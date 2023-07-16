package account

import (
	"errors"
	"github.com/golang/mock/gomock"
	"silver-clean-code/internal/adapter"
	mock_account "silver-clean-code/internal/adapter/account/mock"
	mock_adapter "silver-clean-code/internal/adapter/mock"
	"testing"
)

func TestAccountController_FindByID(t *testing.T) {
	type args struct {
		ctx adapter.ContextServer
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctxMock := mock_adapter.NewMockContextServer(ctrl)
	usecaseMock := mock_account.NewMockUseCase(ctrl)
	gomock.InOrder(
		usecaseMock.EXPECT().GetByID(gomock.Any()).Return(Account{}, nil),
		usecaseMock.EXPECT().GetByID(gomock.Any()).Return(nil, errors.New("error")),
	)
	c := NewAccountController(usecaseMock)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := c.FindByID(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("FindByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
