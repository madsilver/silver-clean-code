package echo

import "github.com/labstack/echo/v4"

type context struct {
	ctx echo.Context
}

func NewEchoContext(c echo.Context) *context {
	return &context{ctx: c}
}

func (s *context) Param(name string) string {
	return s.ctx.Param(name)
}

func (s *context) JSON(code int, i any) error {
	return s.ctx.JSON(code, i)
}

func (s *context) Bind(i interface{}) error {
	return s.ctx.Bind(i)
}
