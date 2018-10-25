// pacote contem toda a regra de negocio da api
package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/bayerlein/red-coins/models"
	"github.com/bayerlein/red-coins/repositories"

	"github.com/patrickmn/go-cache"
)

type TransactionType int

const (
	SELL = 1
	BUY  = 2
)

// Cria uma estrutura de cache para bitcoins com expurgo realizado 1 vez a cada hora
var bitCoinCache = cache.New(1*time.Hour, 1*time.Hour)

// define uma estrutura para o serviço
type BitCoinService struct {
	Repository *repositories.BitCoinTransactionRepository
}

// retorna um ponteiro BitCoinService
func NewBitCoinService() *BitCoinService {
	bitcoinRepository := repositories.NewBitCoinTransactionRepository()
	return &BitCoinService{Repository: bitcoinRepository}
}

// retorna um relatorio por usuário e um texto explicitando o status do processamento
func (service *BitCoinService) GenerateReportByUser(userId int) ([]models.BitCoinTransaction, string) {

	transactions, err := service.Repository.GenerateReportByUserID(userId)

	if err != nil {
		return nil, err.Error()
	}

	return transactions, fmt.Sprintf("Relatorio usuario: %d", userId)
}

// retorna um relatorio por data e um texto explicitando o status do processamento
func (service *BitCoinService) GenerateReportByDate(date string) ([]models.BitCoinTransaction, string) {
	transactions, err := service.Repository.GenerateReportByDate(date)

	if err != nil {
		return nil, err.Error()
	}

	return transactions, fmt.Sprintf("Relatorio data: %s", date)

}

// registra e retorna uma transação de compra de bitcoins, e um texto explicitando o status do processamento
func (service *BitCoinService) BuyBitCoins(amount float64, user models.User) (models.BitCoinTransaction, string) {
	price, err := getBitCoinPrice()
	transaction := models.BitCoinTransaction{}
	transaction.Amount = amount
	transaction.Date = time.Now()
	transaction.PriceUsed = price
	transaction.Total = (amount * price)
	transaction.Type = BUY
	transaction.User = models.User{Id: user.Id}

	err = service.Repository.RegisterTransaction(transaction)

	if err != nil {
		return models.BitCoinTransaction{}, err.Error()
	}
	return transaction, "Compra realizada com sucesso."
}

// registra e retorna uma transação de venda de bitcoins, e um texto explicitando o status do processamento
func (service *BitCoinService) SellBitCoins(amount float64, user models.User) (models.BitCoinTransaction, string) {
	price, err := getBitCoinPrice()
	transaction := models.BitCoinTransaction{}
	transaction.Amount = amount
	transaction.Date = time.Now()
	transaction.PriceUsed = price
	transaction.Total = (amount * price)
	transaction.Type = SELL
	transaction.User = models.User{Id: user.Id}

	err = service.Repository.RegisterTransaction(transaction)

	if err != nil {
		return models.BitCoinTransaction{}, err.Error()
	}

	return transaction, "Venda realizada com sucesso."
}

// função privada que retorna o preço do bitcoin
func getBitCoinPrice() (float64, error) {
	fmt.Println("TENTANDO O CACHE")
	priceCache, found := bitCoinCache.Get("bitcoin_price") // primeira tentativa é buscar o preço que está no cache
	if found {
		fmt.Println(priceCache)
		return priceCache.(float64), nil // se encontrou no cache, então retorna o preço
	} else {
		fmt.Println("NAO ACHOU NO CACHE")
		return getBitCoinPriceByAPI() // caso não tenha encontrado no cache, é chamado um serviço que retorna as informações atuais sobre o bitcoin
	}
}

// função privada que acessa uma api externa de bitcoin e retorna o seu preço
func getBitCoinPriceByAPI() (float64, error) {
	fmt.Println("CHAMANDO API REST PRECO BITCOIN")
	var responseOBJ models.BitCoinResponse
	res, err := http.Get(os.Getenv("BITCOIN_INFO_URL"))

	var bitcoinPrice float64

	if err != nil {
		fmt.Println("Erro: ", err.Error())
		return 0.0, err
	} else {
		data, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(data, &responseOBJ)
		bitcoinPrice = responseOBJ.Data.Quotes.TypeBRL.Price // pega o preço atual do bitcoin
		fmt.Println("PRECO PELA API: ", bitcoinPrice)
		bitCoinCache.Set("bitcoin_price", bitcoinPrice, cache.DefaultExpiration) // registra o preço atual no cache para ser usado numa futura requisição
	}
	return bitcoinPrice, nil
}
