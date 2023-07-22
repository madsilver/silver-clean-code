package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type MysqlDB struct {
	Conn *sql.DB
}

func NewMysqlDB() *MysqlDB {
	dsn := "silver:silver@/silverlabs"
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}

	if err := conn.Ping(); err != nil {
		panic(err.Error())
	}

	conn.SetConnMaxLifetime(time.Minute * 3)
	conn.SetMaxOpenConns(3)
	conn.SetMaxIdleConns(3)

	return &MysqlDB{
		Conn: conn,
	}
}

func (m *MysqlDB) Query(query string, fn func(scan func(dest ...any) error)) error {
	rows, err := m.Conn.Query(query)
	if err != nil {
		return err
	}

	for rows.Next() {
		fn(rows.Scan)
	}

	return nil
}

func (m *MysqlDB) QueryRow(query string, args any, fn func(scan func(dest ...any) error)) error {
	row := m.Conn.QueryRow(query, args)
	if row.Err() != nil {
		return row.Err()
	}

	fn(row.Scan)

	return nil
}

func (m *MysqlDB) Save(query string, args ...any) (any, error) {
	result, err := m.Conn.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
