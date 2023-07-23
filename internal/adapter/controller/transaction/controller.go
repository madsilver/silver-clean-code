package transaction

import (
	"github.com/labstack/gommon/log"
	"net/http"
	"silver-clean-code/internal/adapter"
	"silver-clean-code/internal/adapter/presenter"
	"silver-clean-code/internal/entity"
	"strconv"
)

type UseCase interface {
	GetTransaction(id uint64) (*entity.Transaction, error)
	GetTransactions() ([]*entity.Transaction, error)
	SaveTransaction(transaction *entity.Transaction) error
}

type TransactionController struct {
	usecase UseCase
}

func NewTransactionController(usecase UseCase) *TransactionController {
	return &TransactionController{
		usecase: usecase,
	}
}

// FindTransactionByID godoc
// @Summary Find transaction by ID.
// @Description Find a transaction by the ID.
// @Tags Transaction
// @Accept  json
// @Produce  json
// @Param id path int true "Transaction ID"
// @Success 200 {object} presenter.Transaction
// @Failure 404 {object} presenter.ErrorResponse
// @Failure 500 {object} presenter.ErrorResponse
// @Router /transactions/{id} [get]
func (c *TransactionController) FindTransactionByID(ctx adapter.ContextServer) error {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 8)
	result, err := c.usecase.GetTransaction(id)
	if err != nil {
		log.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, presenter.InternalErrorResponse())
	}

	if result.AccountID == 0 {
		log.Infof("transaction id not found: %v", ctx.Param("id"))
		return ctx.JSON(http.StatusNotFound, presenter.NewErrorResponse("Transaction not found", ""))
	}

	return ctx.JSON(http.StatusOK, presenter.NewTransactionPresenter(result))
}

// FindTransactions godoc
// @Summary List transactions.
// @Description List all transactions.
// @Tags Transaction
// @Accept  json
// @Produce  json
// @Success 200 {object} presenter.Transactions
// @Failure 500 {object} presenter.ErrorResponse
// @Router /transactions [get]
func (c *TransactionController) FindTransactions(ctx adapter.ContextServer) error {
	result, err := c.usecase.GetTransactions()
	if err != nil {
		log.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, presenter.InternalErrorResponse())
	}

	return ctx.JSON(http.StatusOK, presenter.NewTransactionsPresenter(result))
}

// CreateTransaction godoc
// @Summary Create transaction.
// @Description Create new transactions.
// @Tags Transaction
// @Accept json
// @Produce json
// @Param Transaction body presenter.Transaction true " "
// @Success 201 {object} presenter.Transaction
// @Failure 400 {object} presenter.ErrorResponse
// @Failure 500 {object} presenter.ErrorResponse
// @Router /transactions [post]
func (c *TransactionController) CreateTransaction(ctx adapter.ContextServer) error {
	body := presenter.NewTransactionPresenter(nil)
	if err := ctx.Bind(body); err != nil {
		log.Info(err.Error())
		return ctx.JSON(http.StatusBadRequest, presenter.NewErrorResponse("Bad request", err.Error()))
	}
	tran := body.ToEntity()
	err := c.usecase.SaveTransaction(tran)
	if err != nil {
		log.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, presenter.InternalErrorResponse())
	}
	return ctx.JSON(http.StatusCreated, presenter.NewTransactionPresenter(tran))
}
