package controllers

import (
	"net/http"
	"red-coins/models"
	"red-coins/services"
	"red-coins/utils"

	"github.com/go-chi/render"

	"github.com/go-chi/chi"
)

var service = services.NewUserService()

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (user UserController) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", CreateNewUser)

	return router
}

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
