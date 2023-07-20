package db

type DB interface {
	Query(query string, fn func(scan func(dest ...any) error)) error
	QueryRow(query string, args any, fn func(scan func(dest ...any) error)) error
	Save(query string, args any) (any, error)
}
