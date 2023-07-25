package account

import (
	"github.com/labstack/gommon/log"
	"net/http"
	"silver-clean-code/internal/adapter/controller"
	"silver-clean-code/internal/adapter/presenter"
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
// @Success 200 {object} presenter.Account
// @Failure 404 {object} presenter.ErrorResponse
// @Failure 500 {object} presenter.ErrorResponse
// @Router /accounts/{id} [get]
func (c *AccountController) FindAccountByID(ctx controller.ContextServer) error {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 8)
	result, err := c.usecase.GetAccount(id)
	if err != nil {
		log.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, presenter.InternalErrorResponse())
	}

	if result.AccountID == 0 {
		log.Infof("account id not found: %v", ctx.Param("id"))
		return ctx.JSON(http.StatusNotFound, presenter.NewErrorResponse("account not found", ""))
	}

	return ctx.JSON(http.StatusOK, presenter.NewAccountPresenter(result))
}

// FindAccounts godoc
// @Summary List accounts.
// @Description List all accounts.
// @Tags Account
// @Accept  json
// @Produce  json
// @Success 200 {object} presenter.Accounts
// @Failure 500 {object} presenter.ErrorResponse
// @Router /accounts [get]
func (c *AccountController) FindAccounts(ctx controller.ContextServer) error {
	result, err := c.usecase.GetAccounts()
	if err != nil {
		log.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, presenter.InternalErrorResponse())
	}
	return ctx.JSON(http.StatusOK, presenter.NewAccountsPresenter(result))
}

// CreateAccount godoc
// @Summary Create account.
// @Description Create new accounts.
// @Tags Account
// @Accept json
// @Produce json
// @Param Account body presenter.Account true " "
// @Success 201 {object} presenter.Account
// @Failure 400 {object} presenter.ErrorResponse
// @Failure 500 {object} presenter.ErrorResponse
// @Router /accounts [post]
func (c *AccountController) CreateAccount(ctx controller.ContextServer) error {
	body := presenter.NewAccountPresenter(nil)
	if err := ctx.Bind(body); err != nil {
		log.Info(err.Error())
		return ctx.JSON(http.StatusBadRequest, presenter.NewErrorResponse("Bad request", err.Error()))
	}
	acc := body.ToEntity()
	err := c.usecase.SaveAccount(acc)
	if err != nil {
		log.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, presenter.InternalErrorResponse())
	}
	return ctx.JSON(http.StatusCreated, presenter.NewAccountPresenter(acc))
}
