package repositories

import (
	"database/sql"
	"fmt"
)

var server = "localhost"
var userdb = "REDCOINS"
var password = "red_ventures"

type Db struct{}

var instance *Db

func GetDBInstance() *Db {
	if instance == nil {
		instance = &Db{}
	}
	return instance
}

func (db *Db) GetConnectionPool() (*sql.DB, error) {
	fmt.Println("CRIANDO POOL")

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=RedCoins",
		server, userdb, password)
	return sql.Open("mssql", connString)
}
