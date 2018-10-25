// Pacote contem todos os models compartilhados na aplicação
package models

import "time"

type User struct {
	Password     string    `json:"-"`
	Id           int       `json:"-"`
	Email        string    `json:"email"`
	FullName     string    `json:"full_name"`
	DateOfBirth  string    `json:"date_of_birth"`
	RegisterDate time.Time `json:"register_date"`
}
