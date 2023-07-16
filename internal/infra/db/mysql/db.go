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
	//defer conn.Close()

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
