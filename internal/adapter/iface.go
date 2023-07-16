package adapter

type ContextServer interface {
	Param(name string) string
	JSON(code int, i any) error
	Bind(i interface{}) error
}

type DB interface {
	CreateTable(tableName string) error
	FindByID(table string, id any) (any, error)
	Add(tableName string, value any) error
}
