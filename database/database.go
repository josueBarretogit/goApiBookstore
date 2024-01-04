package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type IDb interface {
	Migrate()
	Connect()
	Disconnect()
}

type IRepository interface {
	Create(model interface{}) (error, interface{})
}

type GORMRepositoryService struct {
	repository IRepository
}

func ConnectToDB() (dbInstance *gorm.DB, err error) {
	var error error

	dsn := os.Getenv("DB_URI")
	DB, error := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if error != nil {
		return nil, error

	}
	return DB, nil
}
