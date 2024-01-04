package database

type GORMDbRepository struct{}

func (gormRepo GORMDbRepository) Create(model interface{}) (err error, createModel interface{}) {

	dbInstance, err := ConnectToDB()

	if err != nil {
		return err, nil
	}

	return nil, dbInstance.Create(model)
}
