package account

import (
	"github.com/labstack/gommon/log"
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

// FindAccountByID godoc
// @Summary Find account by ID.
// @Description Find an account by the ID.
// @Tags Account
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} account.Account
// @Failure 404
// @Failure 500
// @Router /accounts/{id} [get]
func (c *AccountController) FindAccountByID(ctx adapter.ContextServer) error {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 8)
	result, err := c.usecase.GetAccount(id)
	if err != nil {
		log.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "internal server error",
		})
	}

	if result.AccountID == 0 {
		log.Infof("account id not found: %v", ctx.Param("id"))
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"message": "account not found",
		})
	}

	return ctx.JSON(http.StatusOK, account.ToPresenter(result))
}

// FindAccounts godoc
// @Summary List accounts.
// @Description List all accounts.
// @Tags Account
// @Accept  json
// @Produce  json
// @Success 200 {object} account.Accounts
// @Failure 500
// @Router /accounts [get]
func (c *AccountController) FindAccounts(ctx adapter.ContextServer) error {
	result, err := c.usecase.GetAccounts()
	if err != nil {
		log.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "internal error",
		})
	}
	return ctx.JSON(http.StatusOK, account.ToPresenters(result))
}

// CreateAccount godoc
// @Summary Create account.
// @Description Create new accounts.
// @Tags Account
// @Accept json
// @Produce json
// @Param Account body account.Account true " "
// @Success 201 {object} account.Account
// @Failure 400
// @Failure 500
// @Router /accounts [post]
func (c *AccountController) CreateAccount(ctx adapter.ContextServer) error {
	body := &account.Account{}
	if err := ctx.Bind(body); err != nil {
		log.Info(err.Error())
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "bad request",
		})
	}
	acc := account.ToEntity(body)
	err := c.usecase.SaveAccount(acc)
	if err != nil {
		log.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "internal error",
		})
	}
	return ctx.JSON(http.StatusCreated, account.ToPresenter(acc))
}
