package models

type BitCoinResponse struct {
	Data BitCoinResponseData `json:"data"`
}

type BitCoinResponseData struct {
	Quotes BitCoinResponseBRL `json:"quotes"`
}

type BitCoinResponseBRL struct {
	TypeBRL BitCoinInfomations `json:"BRL"`
}

type BitCoinInfomations struct {
	Price float64 `json:"price"`
}
