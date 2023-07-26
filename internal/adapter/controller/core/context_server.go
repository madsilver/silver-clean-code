package core

type ContextServer interface {
	Param(name string) string
	JSON(code int, i any) error
	Bind(i interface{}) error
}

type HandlerFunc func(ctx ContextServer) error
