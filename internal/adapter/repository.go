package adapter

type Repository interface {
	FindByID(table string, id any) (any, error)
}
