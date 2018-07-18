package config

import (
	"os"
	"fmt"
)

var (
	username string
	password string
	port string
	dbName string
)

const (
	POSTGRE_USERNAME_KEY = "PGUSERNAME"
	POSTGRE_PASSWORD_KEY = "PGPASSWORD"
	POSTGRE_PORT_KEY = "PGPORT"
	POSTGRE_DBNAME_KEY = "DBNAME"
)

func ConnectionString() string {
	username, password, port, dbName := getConfig()
	return fmt.Sprintf("user=%s password=%s port=%s dbname=%s sslmode=disable", username, password, port, dbName) 
}

func getConfig() (string, string, string, string) {
	username = os.Getenv(POSTGRE_USERNAME_KEY)
	password = os.Getenv(POSTGRE_PASSWORD_KEY)
	port = os.Getenv(POSTGRE_PORT_KEY)
	dbName = os.Getenv(POSTGRE_DBNAME_KEY)
	return username, password, port, dbName
}
