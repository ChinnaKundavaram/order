package repository

import (
	"order/models"

	"github.com/go-chassis/openlog"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DbClient     *gorm.DB
	DatabaseName string
}

func (ur *OrderRepository) CreateOrder(order *models.Order) (*models.Order, error) {
	openlog.Debug("Creating order")

	// if ur.DbClient.Where(models.Order{Id: order.Id}).Take(&models.Order{}).Error == nil {
	// 	return nil, fmt.Errorf("A order already exists with the given id: %v", order.Id)
	// }
	result := ur.DbClient.Create(order)

	if result.Error != nil {

		openlog.Error("Error occured while Creating user errors [" + result.Error.Error() + "]")
		return nil, result.Error
	}
	//firstName := user.FirstName

	return order, nil

}

func (ur *OrderRepository) UpdateOrder(Id string, s map[string]interface{}) (*models.Order, error) {
	openlog.Debug("Updating order with id")
	order := &models.Order{}

	result := ur.DbClient.Where(models.Order{ID: order.ID}).Take(&models.Order{}).Update("status", s["status"])
	if result.Error != nil {
		openlog.Error(result.Error.Error())
		openlog.Error("Error occured while updating order")
		return nil, result.Error
	}
	return order, nil
}

func (ur *OrderRepository) DeleteOrder(Id string) error {
	openlog.Debug("Deleting order with Id")
	order := &models.Order{}
	result := ur.DbClient.Where(models.Order{ID: order.ID}).Take(&models.Order{}).Delete(order)
	if result.Error != nil {
		openlog.Error(result.Error.Error())
		openlog.Error("Error occured while deleting order")
		return result.Error
	}
	return nil
}

func (ur *OrderRepository) FetchAll(Id, page, size int, filters map[string]interface{}) ([]models.Order, error) {
	openlog.Debug("Fetching all data")
	result := make([]models.Order, 0)
	d := ur.DbClient.Offset(page).Limit(size).Where("id", Id).Find(&result)
	if d.Error != nil {
		openlog.Error("Error occured while fetching errors [" + d.Error.Error() + "]")
		return nil, d.Error
	}
	return result, nil
}
func (ur *OrderRepository) FetchTop3() ([]models.Order, error) {
	openlog.Debug("Fetching top three orders")
	result := make([]models.Order, 0)
	d := ur.DbClient.Limit(3).Find(&result)
	if d.Error != nil {
		openlog.Error("Error occured while fetching errors [" + d.Error.Error() + "]")
		return nil, d.Error
	}
	return result, nil
}
