package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type IDB interface {
	Migrate()
	Connect()
	Disconnect()
}

type IRepository interface {
	Find(model interface{}) error
	Create(model interface{}) error
	Update(modelToUpdate interface{}, id uint) error
}

type Database struct {
	Db IDB
}

var DB *gorm.DB

func ConnectToDB() (err error) {
	var error error

	dsn := os.Getenv("DB_URI")
	DB, error = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if error != nil {
		return error
	}

	return nil
}
