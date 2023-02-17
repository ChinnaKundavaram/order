/**
 * Server starts here
 * for reference
 * - openlog - https://github.com/go-chassis/openlog/blob/master/openlog.go
 * - Open architecture - https://medium.com/hackernoon/golang-clean-archithecture-efd6d7c43047
**/

package main

import (
	//_ "go_chassis_template/chassisHandlers"
	"order/database"
	"order/models"
	OrderRepo "order/repository"
	resource "order/resource"
	Orderservice "order/service"

	"github.com/go-chassis/go-archaius"
	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/openlog"
)

func getService() *Orderservice.OrderService {
	repo := OrderRepo.OrderRepository{DbClient: database.GetClient(), DatabaseName: "test_db"}

	return &Orderservice.OrderService{Repo: &repo}
}

func main() {

	temp_resource := resource.OrderResource{}
	chassis.RegisterSchema("rest", &temp_resource) // registration of resources
	if err := chassis.Init(); err != nil {
		openlog.Fatal("Init failed." + err.Error())
		return
	}

	// Add database configurations to archaius
	if err := archaius.AddFile("./conf/database.yaml"); err != nil {
		openlog.Error("add props configurations failed." + err.Error())
		return
	}
	// Add schema paths configurations to archaius
	if err := archaius.AddFile("./conf/payloadSchemas.yaml"); err != nil {
		openlog.Error("add props configurations failed." + err.Error())
		return
	}

	// Server will not start if error occurs.
	if err := database.Connect(); err != nil {
		openlog.Fatal("Error occured while connecting to database")
		return
	}
	models.InitializeModels()
	// Inject service into resource
	temp_resource.Inject(getService())
	// websocket.CreateUpgrader(func(r *http.Request) bool {
	// 	return true
	// })
	chassis.Run()
}
