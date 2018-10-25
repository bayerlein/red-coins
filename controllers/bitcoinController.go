//Pacote que contem todos os mapeamentos de rotas da API
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

// Define uma estrutura para as rotas relacionadas ao BitCoin
type BitCoinController struct{}

// Instancia uma service de Bitcoin, onde contem as regras de negocio
var bitCoinService = services.NewBitCoinService()

// Cria um ponteiro para a autenticação usando JWT
var tokenAuth *jwtauth.JWTAuth

// Função de inicialização usada para configurar o ponteiro de autenticação JWT
func init() {

	// Define o token com o algoritmo de criptografia HS256 usando como chave secreta a string "My secret"
	tokenAuth = jwtauth.New("HS256", []byte("My secret"), nil)
}

// Função que retorna um ponteiro do tipo BitCoinController
func NewBitCoinController() *BitCoinController {
	return &BitCoinController{}
}

// Função que define os mapeamentos das rotas de bitcoin
func (bitCoin BitCoinController) Routes() *chi.Mux {

	router := chi.NewRouter()
	router.Use(jwtauth.Verifier(tokenAuth)) // define o padrão esperado que o token JWT esteja
	router.Use(jwtauth.Authenticator)       // define que as rotas usarão autenticador JWT
	router.Get("/sell/{amount}", SellBitCoins)
	router.Get("/buy/{amount}", BuyBitCoins)
	router.Get("/reports/byuser/{user_id}", GenerateReportByUser)
	router.Get("/reports/byday/{date}", GenerateReportByDate)

	return router
}

// Função gera um relatório de transações de bitcoin, dado algum user_id
func GenerateReportByUser(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "user_id")                           // pega o parametro user_id
	userId, _ := strconv.Atoi(param)                              // converte o user_id para inteiro
	value, message := bitCoinService.GenerateReportByUser(userId) // chama o metodo da camada de serviço que retorna o relatório
	response := utils.CreateResponseObject(value, message)        // cria um objeto que serve como response
	render.JSON(w, r, response)                                   // retorna o response definido para quem chamou a api
}

// Função gera um relatório de transações de bitcoin, dado alguma data
func GenerateReportByDate(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "date")                             // pega o parametro date
	value, message := bitCoinService.GenerateReportByDate(param) // chama o metodo da camada de serviço que retorna o relatorio por data
	response := utils.CreateResponseObject(value, message)       // cria um objeto que serve como response
	render.JSON(w, r, response)                                  // retorna o response definido para quem chamou a ai
}

// Função que registra uma compra de bitcoins
func BuyBitCoins(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context()) // pega os claims contidos no token JWT
	user_id := claims["user_id"].(float64)           // pega o user_id dos claims

	amount := chi.URLParam(r, "amount") // pega o parametro amount(quantidade de bitcoins)

	s, _ := strconv.ParseFloat(amount, 64)                                         // transforma o parametro em float
	value, message := bitCoinService.BuyBitCoins(s, models.User{Id: int(user_id)}) // chama o metodo da camada de serviço que processa a compra
	response := utils.CreateResponseObject(value, message)                         // cria um objeto que serve como response
	render.JSON(w, r, response)                                                    // retorna o response definido para quem chamou a api
}

// Função que registra uma venda de bitcoins
func SellBitCoins(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context()) // pega os claims contidos no token JWT
	user_id := claims["user_id"].(float64)           // pega o user_id dos claims

	amount := chi.URLParam(r, "amount")                                             // pega o parametro amount(quantidade de bitcoins)
	s, _ := strconv.ParseFloat(amount, 64)                                          // tranforma o parametro em float
	value, message := bitCoinService.SellBitCoins(s, models.User{Id: int(user_id)}) // chama o metodo da camada de serviço que processa a venda

	response := utils.CreateResponseObject(value, message) // cria um objeto que serve como response

	render.JSON(w, r, response) // retorna o response definido para quem chamou a api
}
