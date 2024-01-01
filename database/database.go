package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type Db interface {
	Migrate()
	Connect()
	Disconnect()
}

func ConnectToDB() *gorm.DB {
	var error error

	dsn := os.Getenv("DB_URI")
	DB, error := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if error != nil {
		log.Fatal("Failed to connect to db")
	}
	return DB
}

var DB *gorm.DB = ConnectToDB()
