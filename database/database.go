package database

import (
	"github.com/go-chassis/go-archaius"
	"github.com/go-chassis/openlog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


// type global	2

var instance *gorm.DB

// Connects to database
func Connect() error {
	host := archaius.GetString("database.host", "localhost")
	user := archaius.GetString("database.user", "root")
	dbname := archaius.GetString("database.dbname", "test_db")
	password := archaius.GetString("database.password", "root")
	port := archaius.GetString("database.port", "5432")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // <-- thread safe\

	if err != nil {
		openlog.Error("error occured while connecting database")
		return err
	}

	instance = db
	return nil
}

// Provides the instance of the database
func GetClient() *gorm.DB {
	return instance
}
