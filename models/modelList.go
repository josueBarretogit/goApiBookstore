package models

import "api/bookstoreApi/controllers"

type ModelType struct {
	name       string
	controller controllers.IController
}
