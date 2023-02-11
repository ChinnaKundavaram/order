package services

import (
	common "order/common"
	"order/models"
)

type OrderServiceInterface interface {
	CreateOrder(*models.Order) common.HTTPResponse
	UpdateOrder(Id string, s map[string]interface{}) common.HTTPResponse
	DeleteOrder(string) common.HTTPResponse
	FetchAll(Id, pageno, size int, filters map[string]interface{}) common.HTTPResponse
}
