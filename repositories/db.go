package repositories

import (
	"database/sql"
	"fmt"
	"os"
)

var server = os.Getenv("SQL_SERVER_IP")
var userdb = os.Getenv("SQL_SERVER_USER")
var password = os.Getenv("SQL_SERVER_PASS")

type Db struct{}

var instance *Db

func GetDBInstance() *Db {
	if instance == nil {
		instance = &Db{}
	}
	return instance
}

func (db *Db) GetConnectionPool() (*sql.DB, error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=RedCoins",
		server, userdb, password)
	return sql.Open("mssql", connString)
}
