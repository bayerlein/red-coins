//Pacote que contem todos os mapeamentos de rotas da API
package controllers

import (
	"net/http"

	"github.com/bayerlein/red-coins/services"

	"github.com/bayerlein/red-coins/models"
	"github.com/bayerlein/red-coins/utils"
	"github.com/go-chi/render"

	"github.com/go-chi/chi"
)

// Instancia uma service de usuário, onde contem as regras de negocio
var service = services.NewUserService()

// Define uma estrutura para as rotas relacionadas à usuário
type UserController struct{}

// Função que retorna um ponteiro do tipo UserController
func NewUserController() *UserController {
	return &UserController{}
}

// Define o mapeamento das rotas que serão expostas
func (user UserController) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", CreateNewUser)

	return router
}

// Processa a requisição para cadastrar um novo usuário
func CreateNewUser(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	fullName := r.FormValue("full_name")
	password := r.FormValue("password")
	dateOfBirth := r.FormValue("date_of_birth")
	email := r.FormValue("email")

	user := models.User{
		FullName:    fullName,
		Email:       email,
		Password:    password,
		DateOfBirth: dateOfBirth,
	}
	usr := service.CreateNewUser(user)
	response := utils.CreateResponseObject(usr, "Usuário cadastrado com sucesso.")
	render.JSON(w, r, response)
}
