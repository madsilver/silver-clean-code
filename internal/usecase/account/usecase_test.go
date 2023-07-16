package account

import (
	"errors"
	"github.com/golang/mock/gomock"
	"reflect"
	"silver-clean-code/internal/entity"
	mockaccount "silver-clean-code/internal/usecase/account/mock"
	"testing"
)

func Test_accountUseCase_GetByID(t *testing.T) {
	tests := []struct {
		name    string
		id      uint64
		want    *entity.Account
		wantErr bool
	}{
		{
			name:    "should return an account",
			id:      1,
			want:    &entity.Account{AccountID: 1},
			wantErr: false,
		},
		{
			name:    "should return an error",
			id:      1,
			want:    nil,
			wantErr: true,
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mockaccount.NewMockRepository(ctrl)
	gomock.InOrder(
		repository.EXPECT().FindByID(gomock.Any()).Return(&entity.Account{AccountID: 1}, nil),
		repository.EXPECT().FindByID(gomock.Any()).Return(nil, errors.New("error")),
	)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewAccountUseCase(repository)
			got, err := a.GetByID(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
