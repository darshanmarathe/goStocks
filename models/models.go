package models

import (
	"time"
)

type Stock struct {
	Symbol       string    `header:"Symbol"`
	BuyPrice     float64   `header:"BuyPrice"`
	BuyValue     float64   `header:"BuyValue"`
	BuyDate      time.Time `header:"BuyDate"`
	Qty          int       `header:"Qty"`
	SellPrice    float64   `header:"SellPrice"`
	SellValue    float64   `header:"SellValue"`
	SellDate     time.Time `header:"SellDate"`
	GrossProfite float64   `header:"GrossProfite"`
	ProfitBooked float64   `header:"ProfitBooked"`
	IsIntraDay   bool      `header:"IsIntraDay"`
	IsSold       bool      `header:"IsSold"`
}

type StocksType []Stock

func (p StocksType) Len() int {
	return len(p)
}

func (p StocksType) Less(i, j int) bool {
	return p[i].BuyDate.Before(p[j].BuyDate)
}

func (p StocksType) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
