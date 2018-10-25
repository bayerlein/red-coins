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

// compara alguma senha com algum hash e retorna true caso sejam equivalentes
func ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// gera um hash do password informado
func EncryptPassword(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)

	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
