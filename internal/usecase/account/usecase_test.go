package account

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"silver-clean-code/internal/entity"
	mockaccount "silver-clean-code/internal/usecase/account/mock"
	"testing"
)

func Test_accountUseCase_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mockaccount.NewMockRepository(ctrl)
	repository.EXPECT().FindByID(gomock.Any()).Return(&entity.Account{AccountID: 1}, nil)
	a := NewAccountUseCase(repository)

	account, err := a.GetAccount(1)

	assert.NotNil(t, account)
	assert.Nil(t, err)
}

func TestAccountUseCase_GetAccounts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mockaccount.NewMockRepository(ctrl)
	repository.EXPECT().FindAll().Return([]*entity.Account{{AccountID: 1}}, nil)
	a := NewAccountUseCase(repository)

	_, err := a.GetAccounts()

	assert.Nil(t, err)
}

func TestAccountUseCase_SaveAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mockaccount.NewMockRepository(ctrl)
	repository.EXPECT().Save(gomock.Any()).Return(nil)
	a := NewAccountUseCase(repository)

	err := a.SaveAccount(&entity.Account{})

	assert.Nil(t, err)
}
