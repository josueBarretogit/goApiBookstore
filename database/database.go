package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IDB interface {
	Migrate()
	Connect()
	Disconnect()
}

type IRepository interface {
	Find(model interface{}) error
	Create(model interface{}) error
	Update(modelToUpdate interface{}, data interface{}) error
	FindOneBy(modelToFind interface{}, conditions ...interface{}) error
	Delete(modelToDelete interface{}, conditions ...interface{}) error
}

type Database struct {
	Db IDB
}

var DB *gorm.DB

var Pg *pgxpool.Pool

func ConnectToDB() (err error) {
	var error error

	dsn := os.Getenv("DB_URI")
	DB, error = gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if error != nil {
		return error
	}

	Pg, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		return error
	}

	return nil
}
