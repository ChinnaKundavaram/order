package services

import (
	"fmt"
	common "order/common"
	"order/models"
	repository "order/repository"

	"github.com/go-chassis/openlog"
)

type OrderService struct {
	Repo repository.OrderRepositoryInterface
}

func (ur *OrderService) CreateOrder(order *models.Order) common.HTTPResponse {
	openlog.Debug("Creating order")

	res, err := ur.Repo.CreateOrder(order)
	if err != nil {

		openlog.Error("Error occured while creating order [" + err.Error() + "]")
		return common.HTTPResponse{Msg: "Error occured while creating order", Status: 400}

	}

	return common.HTTPResponse{Msg: "Creating order successfully", Data: res, Status: 201}
}

func (ur *OrderService) UpdateOrder(Id string, s map[string]interface{}) common.HTTPResponse {

	res, err := ur.Repo.UpdateOrder(Id, s)

	if err != nil {
		openlog.Error("Error occured while Updating order")
		return common.HTTPResponse{Msg: "Error occured while Updating order", Status: 500}
	}
	fmt.Println(res)

	return common.HTTPResponse{Msg: "Order Updated successfully", Data: res, Status: 200}

}

func (ur *OrderService) DeleteOrder(Id string) common.HTTPResponse {
	err := ur.Repo.DeleteOrder(Id)
	if err != nil {
		openlog.Error("Error occured while Deleting order")
		return common.HTTPResponse{Msg: "Error occured while Deleting order", Status: 500}
	}
	return common.HTTPResponse{Msg: "Order deleted successfully", Status: 200}

}

func (ts *OrderService) FetchAll(Id, page, size int, filters map[string]interface{}) common.HTTPResponse {
	res, err := ts.Repo.FetchAll(Id, page, size, filters)
	if err != nil {
		openlog.Error("Error occured while fetching data")
		return common.HTTPResponse{Msg: "Error occured while fetching data", Status: 500}
	}
	return common.HTTPResponse{Msg: "Fetched orders successfully", Data: res, Status: 200}
}
func (ts *OrderService) FetchTop3() common.HTTPResponse {
	res, err := ts.Repo.FetchTop3()
	if err != nil {
		openlog.Error("Error occured while fetching data")
		return common.HTTPResponse{Msg: "Error occured while fetching data", Status: 500}
	}
	return common.HTTPResponse{Msg: "Fetched orders successfully", Data: res, Status: 200}
}
