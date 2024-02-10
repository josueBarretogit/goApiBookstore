package database

type GORMDbRepository struct{

}

func (gormRepo *GORMDbRepository) Create(model interface{}) (err error) {
	errCreation := DB.Create(&model)
	return errCreation.Error
}

func (gormRepo *GORMDbRepository) Find(model interface{}) (err error) {
	findErr := DB.Find(model)
	return findErr.Error
}

func (gormRepo *GORMDbRepository) FindOneBy(modelToFind interface{}, conditions ...interface{}) error {
	findErr := DB.First(modelToFind, conditions...)
	return findErr.Error
}

func (gormRepo *GORMDbRepository) Update(modelToUpdate interface{}, data interface{}) (err error) {
	errUpdate := DB.Model(modelToUpdate).Updates(data)
	return errUpdate.Error
}

func (gormRepo *GORMDbRepository) Delete(modelToDelete interface{}, conditions ...interface{}) (err error) {
	errDelete := DB.Unscoped().Delete(modelToDelete, conditions)
	return errDelete.Error
}
