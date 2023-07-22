package gin

import (
	"github.com/gin-gonic/gin"
)

type context struct {
	ctx *gin.Context
}

func NewContext(c *gin.Context) *context {
	return &context{ctx: c}
}

func (s *context) Param(name string) string {
	return s.ctx.Param(name)
}

func (s *context) JSON(code int, i any) error {
	s.ctx.JSON(code, i)
	return nil
}

func (s *context) Bind(i interface{}) error {
	return s.ctx.Bind(i)
}
