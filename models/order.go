package models

import gorm "gorm.io/gorm"

type Order struct {
	gorm.Model
	Id           uint   `json:"id"`
	Status       string `json:"status"`
	CurrencyUnit string `json:"currenyunit"`
	items        []item `json:"items"`
}

type item struct {
	Id          uint   `json:"id"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Quantity    string `json:"quantity"`
}
