// pacote contem os acessos realizados à base de dados
package repositories

import (
	"database/sql"
	"fmt"
	"os"
)

// variaveis para conexão
var server = os.Getenv("SQL_SERVER_IP")
var userdb = os.Getenv("SQL_SERVER_USER")
var password = os.Getenv("SQL_SERVER_PASS")

// define uma estrutura para conexões à base
type Db struct{}

var instance *Db

// função que define como singleton, retornando uma instancia que pode já existir
func GetDBInstance() *Db {
	if instance == nil {
		instance = &Db{}
	}
	return instance
}

// retorna(ou cria) um pool de conexão para sql server
func (db *Db) GetConnectionPool() (*sql.DB, error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=RedCoins",
		server, userdb, password)
	return sql.Open("mssql", connString)
}
