package controller

import (
	"net/http"
	"silver-clean-code/internal/adapter"
	"silver-clean-code/internal/adapter/account/presenter"
	"silver-clean-code/internal/entity"
	"strconv"
)

type UseCase interface {
	GetAccount(id uint64) (*entity.Account, error)
	GetAccounts() ([]*entity.Account, error)
}

type AccountController struct {
	usecase UseCase
}

func NewAccountController(usecase UseCase) *AccountController {
	return &AccountController{
		usecase: usecase,
	}
}

func (c *AccountController) FindByID(ctx adapter.ContextServer) error {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 8)
	result, err := c.usecase.GetAccount(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"error": "not found",
		})
	}
	return ctx.JSON(http.StatusOK, presenter.EntityToPresenter(result))
}

func (c *AccountController) FindAll(ctx adapter.ContextServer) error {
	result, err := c.usecase.GetAccounts()
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"error": "not found",
		})
	}
	return ctx.JSON(http.StatusOK, presenter.EntitiesToPresenters(result))
}
