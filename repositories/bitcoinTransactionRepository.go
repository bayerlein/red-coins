package repositories

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bayerlein/red-coins/models"

	_ "github.com/denisenkom/go-mssqldb"
)

type BitCoinTransactionRepository struct{}

func NewBitCoinTransactionRepository() *BitCoinTransactionRepository {
	return &BitCoinTransactionRepository{}
}

func (repository *BitCoinTransactionRepository) RegisterTransaction(transaction models.BitCoinTransaction) {

	var server = "localhost"
	var userdb = "REDCOINS"
	var password = "red_ventures"

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=RedCoins",
		server, userdb, password)

	condb, errdb := sql.Open("mssql", connString)
	if errdb != nil {
		fmt.Println(" Error open db:", errdb.Error())
	}
	tx, err := condb.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare("INSERT INTO bitcoin_transaction(amount, total_value, price_used, transaction_date, transaction_type, user_id) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, errq := stmt.Exec(transaction.Amount, transaction.Total, transaction.PriceUsed, transaction.Date, transaction.Type, transaction.User.Id)
	if errq != nil {
		log.Fatal(errq)
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

}
