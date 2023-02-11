package repository

import (
	//"go_chassis_template/common"
	"order/models"
)

type OrderRepositoryInterface interface {
	CreateOrder(Order *models.Order) (*models.Order, error)
	UpdateOrder(string, map[string]interface{}) (*models.Order, error)
	DeleteOrder(string) error
	FetchAll(OrderId, page, size int, filters map[string]interface{}) ([]models.Order, error)
}
