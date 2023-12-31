package env

import (
	"os"
	"strconv"
)

const (
	EchoServer       = "echo"
	GinServer        = "gin"
	Port             = "8000"
	MysqlServer      = "mysql"
	MysqlUser        = "silver"
	MysqlPassword    = "silver"
	MysqlDatabase    = "silverlabs"
	MysqlHost        = "localhost"
	MysqlPort        = "3306"
	PostgresServer   = "postgres"
	PostgresUser     = "silver"
	PostgresPassword = "silver"
	PostgresDatabase = "silver"
	PostgresHost     = "localhost"
	PostgresPort     = "5432"
)

func GetString(key string, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	return val
}

func GetInt(key string, defaultValue int) int {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	intVal, err := strconv.Atoi(val)
	if err != nil {
		return defaultValue
	}
	return intVal
}
