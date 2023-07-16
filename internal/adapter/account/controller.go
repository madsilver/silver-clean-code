package account

import (
	"net/http"
	"silver-clean-code/internal/adapter"
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
