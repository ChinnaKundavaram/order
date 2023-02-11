package models

import (
	"order/database"

	"github.com/go-chassis/openlog"
)

// create the schema
func InitializeModels() {
	openlog.Info("Initializing the models")
	db := database.GetClient()
	err := db.AutoMigrate(
		&Order{},
	)
	if err != nil {
		openlog.Error("Error occured while initializng the models [" + err.Error() + "]")
		panic(err)
	}
}
