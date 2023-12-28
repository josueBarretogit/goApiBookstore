package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var error error

	dsn := os.Getenv("DB_URI")
	DB, error := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if error != nil {
		log.Fatal("Failed to connect to db")
	}
	println(DB.NamingStrategy)
}
