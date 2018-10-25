package repositories

import (
	"database/sql"
	"fmt"

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

func (repository *BitCoinTransactionRepository) GenerateReportByDate(date string) ([]models.BitCoinTransaction, error) {
	transactions := make([]models.BitCoinTransaction, 1)
	tsql := fmt.Sprintf("select * from bitcoin_transaction tb where datediff(day, tb.transaction_date, '%s') = 0", date)
	rows, err := repository.Db.Query(tsql)
	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return nil, error
	}
	defer rows.Close()
	for rows.Next() {
		transaction := models.BitCoinTransaction{}

		err := rows.Scan(&transaction.Id, &transaction.Amount, &transaction.Total, &transaction.PriceUsed, &transaction.Date, &transaction.Type, &transaction.User.Id)

		transactions = append(transactions, transaction)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return nil, error
		}
		fmt.Println(transactions)
	}

	return transactions, nil
}

func (repository *BitCoinTransactionRepository) GenerateReportByUserID(userId int) ([]models.BitCoinTransaction, error) {
	transactions := make([]models.BitCoinTransaction, 1)
	tsql := fmt.Sprintf("SELECT id, amount, total_value, price_used, transaction_date, transaction_type, user_id FROM bitcoin_transaction WHERE user_id = %d", userId)
	rows, err := repository.Db.Query(tsql)
	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return nil, error
	}
	defer rows.Close()
	for rows.Next() {
		transaction := models.BitCoinTransaction{}

		err := rows.Scan(&transaction.Id, &transaction.Amount, &transaction.Total, &transaction.PriceUsed, &transaction.Date, &transaction.Type, &transaction.User.Id)

		transactions = append(transactions, transaction)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return nil, error
		}
		fmt.Println(transactions)
	}

	return transactions, nil
}

func (repository *BitCoinTransactionRepository) RegisterTransaction(transaction models.BitCoinTransaction) error {

	if repository.ErrDB != nil {
		fmt.Println(" Error open db:", repository.ErrDB.Error())
		return error
	}
	tx, err := repository.Db.Begin()
	if err != nil {
		fmt.Println(" Error open db:", err.Error())
		return error
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare("INSERT INTO bitcoin_transaction(amount, total_value, price_used, transaction_date, transaction_type, user_id) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println(" Error open db:", err.Error())
		return error
	}
	defer stmt.Close()

	_, errq := stmt.Exec(transaction.Amount, transaction.Total, transaction.PriceUsed, transaction.Date, transaction.Type, transaction.User.Id)
	if errq != nil {
		fmt.Println(" Error open db:", errq.Error())
		return error
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println(" Error open db:", err.Error())
		return error
	}

	return nil

}
