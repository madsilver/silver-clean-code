package account

import (
	"net/http"
	"silver-clean-code/internal/adapter"
	"silver-clean-code/internal/adapter/presenter/account"
	"silver-clean-code/internal/entity"
	"strconv"
)

type UseCase interface {
	GetAccount(id uint64) (*entity.Account, error)
	GetAccounts() ([]*entity.Account, error)
	SaveAccount(account *entity.Account) error
}

type AccountController struct {
	usecase UseCase
}

func NewAccountController(usecase UseCase) *AccountController {
	return &AccountController{
		usecase: usecase,
	}
}

func (c *AccountController) FindAccountByID(ctx adapter.ContextServer) error {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 8)
	result, err := c.usecase.GetAccount(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "internal server error",
		})
	}

	if result.AccountID == 0 {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"message": "account not found",
		})
	}

	return ctx.JSON(http.StatusOK, account.ToPresenter(result))
}

func (c *AccountController) FindAccounts(ctx adapter.ContextServer) error {
	result, err := c.usecase.GetAccounts()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "internal error",
		})
	}
	return ctx.JSON(http.StatusOK, account.ToPresenters(result))
}

func (c *AccountController) CreateAccount(ctx adapter.ContextServer) error {
	body := &account.Account{}
	if err := ctx.Bind(body); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "bad request",
		})
	}
	acc := account.ToEntity(body)
	err := c.usecase.SaveAccount(acc)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "internal error",
		})
	}
	return ctx.JSON(http.StatusOK, account.ToPresenter(acc))
}
