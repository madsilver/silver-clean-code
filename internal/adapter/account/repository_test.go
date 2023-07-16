package account

import (
	"errors"
	"github.com/golang/mock/gomock"
	"reflect"
	mock_adapter "silver-clean-code/internal/adapter/mock"
	"silver-clean-code/internal/entity"
	"testing"
)

func TestAccountRepository_FindByID(t *testing.T) {
	tests := []struct {
		name    string
		id      uint64
		want    *entity.Account
		wantErr bool
	}{
		{
			name:    "",
			id:      1,
			want:    &entity.Account{},
			wantErr: false,
		},
		{
			name:    "",
			id:      1,
			want:    nil,
			wantErr: true,
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dbMock := mock_adapter.NewMockDB(ctrl)
	r := NewAccountRepository(dbMock)
	gomock.InOrder(
		dbMock.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(&entity.Account{}, nil),
		dbMock.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(nil, errors.New("error")),
	)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.FindByID(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
