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
// @Success 200 {object} transaction.Transaction
// @Failure 404
// @Failure 500
// @Router /transactions/{id} [get]
func (c *TransactionController) FindTransactionByID(ctx adapter.ContextServer) error {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 8)
	result, err := c.usecase.GetTransaction(id)
	if err != nil {
		log.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "internal server error",
		})
	}

	if result.AccountID == 0 {
		log.Infof("transaction id not found: %v", ctx.Param("id"))
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"message": "transaction not found",
		})
	}

	return ctx.JSON(http.StatusOK, presenter.NewTransactionPresenter(result))
}

// FindTransactions godoc
// @Summary List transactions.
// @Description List all transactions.
// @Tags Transaction
// @Accept  json
// @Produce  json
// @Success 200 {object} transaction.Transactions
// @Failure 500
// @Router /transactions [get]
func (c *TransactionController) FindTransactions(ctx adapter.ContextServer) error {
	result, err := c.usecase.GetTransactions()
	if err != nil {
		log.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "internal server error",
		})
	}

	return ctx.JSON(http.StatusOK, presenter.NewTransactionsPresenter(result))
}

// CreateTransaction godoc
// @Summary Create transaction.
// @Description Create new transactions.
// @Tags Transaction
// @Accept json
// @Produce json
// @Param Transaction body transaction.Transaction true " "
// @Success 201 {object} transaction.Transaction
// @Failure 400
// @Failure 500
// @Router /transactions [post]
func (c *TransactionController) CreateTransaction(ctx adapter.ContextServer) error {
	body := presenter.NewTransactionPresenter(nil)
	if err := ctx.Bind(body); err != nil {
		log.Info(err.Error())
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "bad request",
		})
	}
	tran := body.ToEntity()
	err := c.usecase.SaveTransaction(tran)
	if err != nil {
		log.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "internal error",
		})
	}
	return ctx.JSON(http.StatusCreated, presenter.NewTransactionPresenter(tran))
}
