package controller

type ContextServer interface {
	Param(name string) string
	JSON(code int, i any) error
	Bind(i interface{}) error
}
