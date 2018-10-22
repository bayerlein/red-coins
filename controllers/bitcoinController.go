package controllers

import (
	"net/http"
	"red-coins/models"
	"red-coins/services"
	"red-coins/utils"
	"strconv"

	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"

	"github.com/go-chi/chi"
)

type BitCoinController struct{}

var bitCoinService = services.NewBitCoinService()

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("My secret"), nil)
}

func NewBitCoinController() *BitCoinController {
	return &BitCoinController{}
}

func (bitCoin BitCoinController) Routes() *chi.Mux {

	router := chi.NewRouter()
	router.Use(jwtauth.Verifier(tokenAuth))
	router.Use(jwtauth.Authenticator)
	router.Get("/sell/{amount}", SellBitCoins)
	router.Get("/buy/{amount}", BuyBitCoins)
	router.Get("/reports/{user_name}", GenerateReportByUser)
	router.Get("/reports/{day}", GenerateReportByDay)

	return router
}

func GenerateReportByUser(w http.ResponseWriter, r *http.Request) {
	bitCoinService.GenerateReportByUser()
}
func GenerateReportByDay(w http.ResponseWriter, r *http.Request) {

	bitCoinService.GenerateReportByDay()
}

func BuyBitCoins(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	user_id := claims["user_id"].(float64)

	amount := chi.URLParam(r, "amount")

	s, _ := strconv.ParseFloat(amount, 64)
	value, message := bitCoinService.BuyBitCoins(s, models.User{Id: int(user_id)})
	response := utils.CreateResponseObject(value, message)
	render.JSON(w, r, response)
}

func SellBitCoins(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	user_id := claims["user_id"].(float64)

	amount := chi.URLParam(r, "amount")
	s, _ := strconv.ParseFloat(amount, 64)
	value, message := bitCoinService.SellBitCoins(s, models.User{Id: int(user_id)})

	response := utils.CreateResponseObject(value, message)

	render.JSON(w, r, response)
}
