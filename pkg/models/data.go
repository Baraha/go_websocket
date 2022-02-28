package models

const COLLECTION_CRYPTOCURRENCY = "cryptocurrency"

type Data struct {
	Coin_id  string `json:"coin_id"`
	Rank     string `json:"rank"`
	Symbol   string `json:"symbol"`
	Interval int    `json:"interval"`
	PriceUsd string `json:"priceUsd"`
	IsHandle bool   `json:"isHandle"`
}
