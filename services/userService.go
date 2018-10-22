//Pacote expoem os metodos contendo as regras de negocio
package services

import (
	"red-coins/models"
	"red-coins/repositories"
	"red-coins/utils"
	"time"
)

type UserService struct{}

var repository = repositories.NewUserRepository()

//Retorna um struct 'UserService' vazio
func NewUserService() *UserService {
	return &UserService{}
}

//Recebe um struct 'user' como parametro e faz a persistencia das informações
func (service *UserService) CreateNewUser(user models.User) models.User {
	hash := utils.EncryptPassword([]byte(user.Password))

	user.Password = hash
	user.RegisterDate = time.Now()

	repository.CreateNewUser(user)
	return user
}
