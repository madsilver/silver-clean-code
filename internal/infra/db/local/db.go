package local

import "errors"

type LocalDB struct {
	storage map[string]map[any]any
}

func NewLocalDB() *LocalDB {
	return &LocalDB{
		storage: map[string]map[any]any{},
	}
}

func (l *LocalDB) CreateTable(tableName string) error {
	if l.storage[tableName] != nil {
		return errors.New("table already exists")
	}
	l.storage[tableName] = make(map[any]any)
	return nil
}

func (l *LocalDB) FindByID(tableName string, id any) (any, error) {
	if result, ok := l.storage[tableName][id]; ok {
		return result, nil
	}
	return nil, errors.New("not found")
}

func (l *LocalDB) Add(tableName string, value any) error {
	l.storage[tableName][0] = value
	return nil
}
