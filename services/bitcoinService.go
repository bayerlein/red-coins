package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

var bitCoinCache = cache.New(1*time.Hour, 1*time.Hour)

type BitCoinService struct {
	Repository *repositories.BitCoinTransactionRepository
}

func NewBitCoinService() *BitCoinService {
	bitcoinRepository := repositories.NewBitCoinTransactionRepository()
	return &BitCoinService{Repository: bitcoinRepository}
}

func (service *BitCoinService) GenerateReportByUser(userId int) {
	service.Repository.GenerateReportByUserID(userId)
}

func (service *BitCoinService) GenerateReportByDay() {
}

func (service *BitCoinService) BuyBitCoins(amount float64, user models.User) (models.BitCoinTransaction, string) {
	price := getBitCoinPrice()
	transaction := models.BitCoinTransaction{}
	transaction.Amount = amount
	transaction.Date = time.Now()
	transaction.PriceUsed = price
	transaction.Total = (amount * price)
	transaction.Type = BUY
	transaction.User = models.User{Id: user.Id}

	service.Repository.RegisterTransaction(transaction)

	return transaction, "Compra realizada com sucesso."
}

func (service *BitCoinService) SellBitCoins(amount float64, user models.User) (models.BitCoinTransaction, string) {
	price := getBitCoinPrice()
	transaction := models.BitCoinTransaction{}
	transaction.Amount = amount
	transaction.Date = time.Now()
	transaction.PriceUsed = price
	transaction.Total = (amount * price)
	transaction.Type = SELL
	transaction.User = models.User{Id: user.Id}

	service.Repository.RegisterTransaction(transaction)

	return transaction, "Venda realizada com sucesso."
}

func getBitCoinPrice() float64 {
	fmt.Println("TENTANDO O CACHE")
	priceCache, found := bitCoinCache.Get("bitcoin_price")
	if found {
		fmt.Println(priceCache)
		return priceCache.(float64)
	} else {
		fmt.Println("NAO ACHOU NO CACHE")

		return getBitCoinPriceByAPI()
	}
}
func getBitCoinPriceByAPI() float64 {
	fmt.Println("CHAMANDO API REST PRECO BITCOIN")
	var responseOBJ models.BitCoinResponse
	res, err := http.Get("https://api.coinmarketcap.com/v2/ticker/1/?convert=BRL")

	var bitcoinPrice float64

	if err != nil {
		fmt.Println("Erro: ", err.Error())
	} else {
		data, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(data, &responseOBJ)
		bitcoinPrice = responseOBJ.Data.Quotes.TypeBRL.Price
		fmt.Println("PRECO PELA API: ", bitcoinPrice)
		bitCoinCache.Set("bitcoin_price", bitcoinPrice, cache.DefaultExpiration)
	}
	return bitcoinPrice
}
