package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"silver-clean-code/internal/infra/env"
	"time"
)

type MysqlDB struct {
	Conn *sql.DB
}

func NewMysqlDB() *MysqlDB {
	conn, err := sql.Open("mysql", getDSN())
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

func getDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		env.GetString("MYSQL_USER", env.MysqlUser),
		env.GetString("MYSQL_PASSWORD", env.MysqlPassword),
		env.GetString("MYSQL_HOST", env.MysqlHost),
		env.GetString("MYSQL_PORT", env.MysqlPort),
		env.GetString("MYSQL_DATABASE", env.MysqlDatabase))
}

func (m *MysqlDB) Query(query string, fn func(scan func(dest ...any) error) error) error {
	rows, err := m.Conn.Query(query)
	if err != nil {
		return err
	}

	for rows.Next() {
		err = fn(rows.Scan)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *MysqlDB) QueryRow(query string, args any, fn func(scan func(dest ...any) error) error) error {
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
