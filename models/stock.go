package models

import "time"

type Stock struct {
	StockID   string    `json:"stock_id`
	Symbol    string    `json:"symbol"`
	Name      string    `json:"name"`
	Price     string    `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
