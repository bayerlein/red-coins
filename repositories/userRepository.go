package repositories

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bayerlein/red-coins/models"

	_ "github.com/denisenkom/go-mssqldb"
)

type UserRepository struct {
	Db    *sql.DB
	ErrDB error
}

func NewUserRepository() *UserRepository {
	connection, err := GetDBInstance().GetConnectionPool()
	return &UserRepository{Db: connection, ErrDB: err}
}

func (repository *UserRepository) CreateNewUser(user models.User) {

	if repository.ErrDB != nil {
		fmt.Println(" Error open db:", repository.ErrDB.Error())
	}
	tx, err := repository.Db.Begin()
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
