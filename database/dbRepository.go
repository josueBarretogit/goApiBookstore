package database

type GORMDbRepository struct{}

func (gormRepo *GORMDbRepository) Create(model interface{}) (err error) {
	errCreation := DB.Create(model)
	return errCreation.Error
}

func (gormRepo *GORMDbRepository) Find(model interface{}) (err error) {
	findErr := DB.Find(model)
	return findErr.Error
}

func (gormRepo *GORMDbRepository) Update(dataToUpdate interface{}, id uint) (err error) {
	errCreation := DB.Update(dataToUpdate, id)
	return errCreation.Error
}
