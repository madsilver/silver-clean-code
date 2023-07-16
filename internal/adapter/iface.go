package adapter

type ContextServer interface {
	Param(name string) string
	JSON(code int, i any) error
}

type DB interface {
	CreateTable(tableName string) error
	FindByID(table string, id any) (any, error)
	Add(tableName string, id any, value any) error
}
