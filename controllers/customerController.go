package controllers

import usermodels "api/bookstoreApi/models/userModels"

type CustomerController struct {
	GenericController[usermodels.Customer]
}
