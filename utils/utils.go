//Pacote que contem metodos para utilidade e que são comuns para toda a aplicação
package utils

import (
	"log"

	"github.com/bayerlein/red-coins/models"

	"golang.org/x/crypto/bcrypt"
)

//Função cria e retorna um struct Response
func CreateResponseObject(data interface{}, message string) models.Response {

	return models.Response{Data: data, Message: message}
}
func ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func EncryptPassword(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)

	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
