package account

import (
	"net/http"
	"silver-clean-code/internal/adapter"
	"silver-clean-code/internal/entity"
	"strconv"
)

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
	result, err := c.usecase.GetByID(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"error": "not found",
		})
	}
	return ctx.JSON(http.StatusOK, EntityToPresenter(result))
}

func (c *AccountController) Create(ctx adapter.ContextServer) error {
	account := &entity.Account{}
	if err := ctx.Bind(account); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "bad request",
		})
	}
	err := c.usecase.Create(account)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "account created",
	})
}
