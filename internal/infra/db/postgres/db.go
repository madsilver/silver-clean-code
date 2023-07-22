package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"silver-clean-code/internal/infra/env"
	"time"
)

type PostgresDB struct {
	Conn *sql.DB
}

func NewPostgresDB() *PostgresDB {
	conn, err := sql.Open("postgres", getDSN())
	if err != nil {
		panic(err.Error())
	}

	if err := conn.Ping(); err != nil {
		panic(err.Error())
	}

	conn.SetConnMaxLifetime(time.Minute * 3)
	conn.SetMaxOpenConns(3)
	conn.SetMaxIdleConns(3)

	return &PostgresDB{
		Conn: conn,
	}
}

func getDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.GetString("POSTGRES_DATABASE", env.PostgresHost),
		env.GetString("POSTGRES_DATABASE", env.PostgresPort),
		env.GetString("POSTGRES_USER", env.PostgresUser),
		env.GetString("POSTGRES_PASSWORD", env.PostgresPassword),
		env.GetString("POSTGRES_DATABASE", env.PostgresDatabase))
}

func (m *PostgresDB) Query(query string, fn func(scan func(dest ...any) error) error) error {
	rows, err := m.Conn.Query(query)
	if err != nil {
		return err
	}

	for rows.Next() {
		if err = fn(rows.Scan); err != nil {
			return err
		}
	}

	return nil
}

func (m *PostgresDB) QueryRow(query string, args any, fn func(scan func(dest ...any) error) error) error {
	row := m.Conn.QueryRow(query, args)
	if row.Err() != nil {
		return row.Err()
	}

	if err := fn(row.Scan); err != nil {
		if err.Error() != "sql: no rows in result set" {
			return err
		}
	}

	return nil
}

func (m *PostgresDB) Save(query string, args ...any) (any, error) {
	lastInsertId := int64(0)
	row := m.Conn.QueryRow(query, args...)
	if row.Err() != nil {
		return 0, row.Err()
	}

	err := row.Scan(&lastInsertId)

	return lastInsertId, err
}
