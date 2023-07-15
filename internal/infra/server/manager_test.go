package server

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_accountUseCase(t *testing.T) {
	manager := NewManager(nil)

	assert.NotNil(t, manager)
	assert.NotNil(t, manager.AccountController)
}
