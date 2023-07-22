package transaction

import (
	"net/http"
	"silver-clean-code/internal/adapter"
	"silver-clean-code/internal/adapter/presenter/transaction"
	"silver-clean-code/internal/entity"
	"strconv"
)

type UseCase interface {
	GetTransaction(id uint64) (*entity.Transaction, error)
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

func (c *TransactionController) FindTransactionByID(ctx adapter.ContextServer) error {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 8)
	result, err := c.usecase.GetTransaction(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "internal server error",
		})
	}

	if result.AccountID == 0 {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"message": "transaction not found",
		})
	}

	return ctx.JSON(http.StatusOK, transaction.ToPresenter(result))
}

func (c *TransactionController) CreateTransaction(ctx adapter.ContextServer) error {
	body := &transaction.Transaction{}
	if err := ctx.Bind(body); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "bad request",
		})
	}
	tran := transaction.ToEntity(body)
	err := c.usecase.SaveTransaction(tran)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "internal error",
		})
	}
	return ctx.JSON(http.StatusOK, transaction.ToPresenter(tran))
}
