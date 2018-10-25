// pacote contem os acessos realizados à base de dados
package repositories

import (
	"database/sql"
	"fmt"

	"github.com/bayerlein/red-coins/models"

	_ "github.com/denisenkom/go-mssqldb"
)

// define uma estrutura para o repositorio de bitcoins
type BitCoinTransactionRepository struct {
	Db    *sql.DB
	ErrDB error
}

// retorna um ponteiro para o repositorio
func NewBitCoinTransactionRepository() *BitCoinTransactionRepository {
	connection, err := GetDBInstance().GetConnectionPool() // pega o pool de conexão criado
	return &BitCoinTransactionRepository{Db: connection, ErrDB: err}
}

// retorna um relatorio usando uma data como filtro
func (repository *BitCoinTransactionRepository) GenerateReportByDate(date string) ([]models.BitCoinTransaction, error) {
	transactions := make([]models.BitCoinTransaction, 1)
	tsql := fmt.Sprintf("select * from bitcoin_transaction tb where datediff(day, tb.transaction_date, '%s') = 0", date)
	rows, err := repository.Db.Query(tsql)
	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		transaction := models.BitCoinTransaction{}

		err := rows.Scan(&transaction.Id, &transaction.Amount, &transaction.Total, &transaction.PriceUsed, &transaction.Date, &transaction.Type, &transaction.User.Id)

		transactions = append(transactions, transaction)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return nil, err
		}
		fmt.Println(transactions)
	}

	return transactions, nil
}

// retorna um relatorio usando um user_id como filtro
func (repository *BitCoinTransactionRepository) GenerateReportByUserID(userId int) ([]models.BitCoinTransaction, error) {
	transactions := make([]models.BitCoinTransaction, 1)
	tsql := fmt.Sprintf("SELECT id, amount, total_value, price_used, transaction_date, transaction_type, user_id FROM bitcoin_transaction WHERE user_id = %d", userId)
	rows, err := repository.Db.Query(tsql)
	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		transaction := models.BitCoinTransaction{}

		err := rows.Scan(&transaction.Id, &transaction.Amount, &transaction.Total, &transaction.PriceUsed, &transaction.Date, &transaction.Type, &transaction.User.Id)

		transactions = append(transactions, transaction)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return nil, err
		}
		fmt.Println(transactions)
	}

	return transactions, nil
}

// registra uma transação, podem ser compra ou venda
func (repository *BitCoinTransactionRepository) RegisterTransaction(transaction models.BitCoinTransaction) error {

	if repository.ErrDB != nil {
		fmt.Println(" Error open db:", repository.ErrDB.Error())
		return repository.ErrDB
	}
	tx, err := repository.Db.Begin()
	if err != nil {
		fmt.Println(" Error open db:", err.Error())
		return err
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare("INSERT INTO bitcoin_transaction(amount, total_value, price_used, transaction_date, transaction_type, user_id) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println(" Error open db:", err.Error())
		return err
	}
	defer stmt.Close()

	_, errq := stmt.Exec(transaction.Amount, transaction.Total, transaction.PriceUsed, transaction.Date, transaction.Type, transaction.User.Id)
	if errq != nil {
		fmt.Println(" Error open db:", errq.Error())
		return errq
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println(" Error open db:", err.Error())
		return err
	}

	return nil

}
