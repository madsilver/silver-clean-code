package echo

import "github.com/labstack/echo/v4"

type Context struct {
	ctx echo.Context
}

func NewContext(c echo.Context) *Context {
	return &Context{ctx: c}
}

func (s *Context) Param(name string) string {
	return s.ctx.Param(name)
}

func (s *Context) JSON(code int, i any) error {
	return s.ctx.JSON(code, i)
}

func (s *Context) Bind(i interface{}) error {
	return s.ctx.Bind(i)
}
