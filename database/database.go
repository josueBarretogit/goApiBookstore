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
	Find()
	FindById(id uint)
	Create(model interface{})
	Update(id uint, model interface{})
	Delete(model interface{}, where ...interface{})
}

type DbRepositoryService struct {
	repository IRepository
}

func NewDbRepositoryService(repo IRepository) *DbRepositoryService {
	return &DbRepositoryService{
		repository: repo,
	}
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
