package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Db interface {
	Migrate()
	Connect()
	Disconnect()
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
