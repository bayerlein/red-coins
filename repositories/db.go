package repositories

import (
	"os"
	"database/sql"
	"fmt"
)

var server = "192.168.15.22"
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
	fmt.Println(os.Getenv("TESTE"))
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=RedCoins",
		server, userdb, password)
	return sql.Open("mssql", connString)
}
