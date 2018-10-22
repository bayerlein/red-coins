package models

import (
	"time"
)

type BitCoinTransaction struct {
	Id        int       `json:"-"`
	Amount    float64   `json:"amount"`
	Total     float64   `json:"total"`
	PriceUsed float64   `json:"price_used"`
	Date      time.Time `json:"date"`
	Type      int       `json:"transaction_type"`
	User      User      `json:"-"`
}
