package repositories

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bayerlein/red-coins/models"

	_ "github.com/denisenkom/go-mssqldb"
)

type BitCoinTransactionRepository struct {
	Db    *sql.DB
	ErrDB error
}

func NewBitCoinTransactionRepository() *BitCoinTransactionRepository {
	connection, err := GetDBInstance().GetConnectionPool()
	return &BitCoinTransactionRepository{Db: connection, ErrDB: err}
}

func (repository *BitCoinTransactionRepository) GenerateReportByUserID(userId int) {
	tsql := fmt.Sprintf("SELECT id, amount FROM bitcoin_transaction")
	rows, err := repository.Db.Query(tsql)
	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
	}
	defer rows.Close()
	count := 0
	for rows.Next() {
		var amount float64
		var id int
		err := rows.Scan(&id, &amount)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
		}
		fmt.Printf("ID: %d, Name: %s, Location: %s\n", id, amount)
		count++
	}
}

func (repository *BitCoinTransactionRepository) RegisterTransaction(transaction models.BitCoinTransaction) {

	if repository.ErrDB != nil {
		fmt.Println(" Error open db:", repository.ErrDB.Error())
	}
	tx, err := repository.Db.Begin()
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
