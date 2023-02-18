package models

import (
	gorm "gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID       uint   `json:"id"`
	Status   string `json:"status"`
	Currency string `json:"currency"`
	Item     string `json:"items"`
}
