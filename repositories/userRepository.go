package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"red-coins/models"

	_ "github.com/denisenkom/go-mssqldb"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (repository *UserRepository) CreateNewUser(user models.User) {

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
	stmt, err := tx.Prepare("INSERT INTO users(full_name, password, email, date_of_birth, register_date) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	stmt.Exec(user.FullName, user.Password, user.Email, user.DateOfBirth, user.RegisterDate)
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

}
