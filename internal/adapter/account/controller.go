package account

import (
	"net/http"
	"silver-clean-code/internal/adapter"
	"silver-clean-code/internal/usecase/account"
	"strconv"
)

type AccountController struct {
	usecase *account.AccountUseCase
}

func NewAccountController(usecase *account.AccountUseCase) *AccountController {
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
