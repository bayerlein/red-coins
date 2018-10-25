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

var bitCoinCache = cache.New(1*time.Hour, 1*time.Hour)

type BitCoinService struct {
	Repository *repositories.BitCoinTransactionRepository
}

func NewBitCoinService() *BitCoinService {
	bitcoinRepository := repositories.NewBitCoinTransactionRepository()
	return &BitCoinService{Repository: bitcoinRepository}
}

func (service *BitCoinService) GenerateReportByUser(userId int) ([]models.BitCoinTransaction, string) {

	transactions, err := service.Repository.GenerateReportByUserID(userId)

	if err != nil {
		return nil, err.Error()
	}

	return transactions, fmt.Sprintf("Relatorio usuario: %d", userId)
}

func (service *BitCoinService) GenerateReportByDate(date string) ([]models.BitCoinTransaction, string) {
	transactions, err := service.Repository.GenerateReportByDate(date)

	if err != nil {
		return nil, err.Error()
	}

	return transactions, fmt.Sprintf("Relatorio data: %s", date)

}

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

func getBitCoinPrice() (float64, error) {
	fmt.Println("TENTANDO O CACHE")
	priceCache, found := bitCoinCache.Get("bitcoin_price")
	if found {
		fmt.Println(priceCache)
		return priceCache.(float64), nil
	} else {
		fmt.Println("NAO ACHOU NO CACHE")
		return getBitCoinPriceByAPI()
	}
}
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
		bitcoinPrice = responseOBJ.Data.Quotes.TypeBRL.Price
		fmt.Println("PRECO PELA API: ", bitcoinPrice)
		bitCoinCache.Set("bitcoin_price", bitcoinPrice, cache.DefaultExpiration)
	}
	return bitcoinPrice, nil
}
