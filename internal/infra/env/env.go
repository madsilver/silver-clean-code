package env

import (
	"os"
	"strconv"
)

const (
	EchoServer    = "echo"
	GinServer     = "gin"
	Port          = "8000"
	TableAccounts = "Accounts"
	MysqlUser     = "silver"
	MysqlPassword = "silver"
	MysqlDatabase = "silverlabs"
)

func GetString(key string, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	return defaultValue
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
