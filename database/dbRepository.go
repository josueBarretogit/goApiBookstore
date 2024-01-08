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

func (gormRepo *GORMDbRepository) FindOneBy(modelToFind interface{}, term ...interface{}) error {
	findErr := DB.First(modelToFind, term)
	return findErr.Error
}

func (gormRepo *GORMDbRepository) Update(modelToUpdate interface{}, id uint) (err error) {
	errUpdate := DB.Model(modelToUpdate).Updates(modelToUpdate)
	return errUpdate.Error
}
