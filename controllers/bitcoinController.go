package controllers

import (
	"net/http"
	"strconv"

	"github.com/bayerlein/red-coins/models"
	"github.com/bayerlein/red-coins/services"
	"github.com/bayerlein/red-coins/utils"
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
	router.Get("/reports/byuser/{user_id}", GenerateReportByUser)
	router.Get("/reports/byday/{date}", GenerateReportByDate)

	return router
}

func GenerateReportByUser(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "user_id")
	userId, _ := strconv.Atoi(param)
	value, message := bitCoinService.GenerateReportByUser(userId)
	response := utils.CreateResponseObject(value, message)
	render.JSON(w, r, response)
}
func GenerateReportByDate(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "date")
	value, message := bitCoinService.GenerateReportByDate(param)
	response := utils.CreateResponseObject(value, message)
	render.JSON(w, r, response)
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
