// pacote contem toda a regra de negocio da api
package services

import (
	"time"

	"github.com/bayerlein/red-coins/models"
	"github.com/bayerlein/red-coins/repositories"
	"github.com/bayerlein/red-coins/utils"
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
