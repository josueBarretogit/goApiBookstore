package routes

import (
	"api/bookstoreApi/consts"
	"api/bookstoreApi/controllers"
)

type ModelListFormat struct {
	ModelName  string
	Controller controllers.IController
}

func ModelList() []ModelListFormat {
	modelList := []ModelListFormat{
		{ModelName: consts.RoleModelName, Controller: controllers.NewRoleController()},
		{ModelName: consts.AuthorModelName, Controller: controllers.NewAuthorController()},
		{ModelName: consts.GenreModelName, Controller: controllers.NewGenreController()},
		{ModelName: consts.AccountModelName, Controller: controllers.NewAccountController()},
		{ModelName: consts.CustomerModelName, Controller: controllers.NewCustomerController()},
		{ModelName: consts.PublisherModelName, Controller: controllers.NewPublisherController()},
		{ModelName: consts.BookModelName, Controller: controllers.NewBookController()},
		{ModelName: consts.HardcoverFormatModelName, Controller: controllers.NewHardCoverFormatController()},
		{ModelName: consts.DigitalFormatModelName, Controller: controllers.NewDigitalFormatController()},
		{ModelName: consts.AudioBookFormatModelName, Controller: controllers.NewAudioBookFormatController()},
		{ModelName: consts.OrderModelName, Controller: controllers.NewOrderController()},
		{ModelName: consts.OrderDetailsModelName, Controller: controllers.NewOrderDetailsController()},
		{ModelName: consts.ReviewModelName, Controller: controllers.NewReviewController()},
		{ModelName: consts.PaymentModelName, Controller: controllers.NewPaymentController()},
		{ModelName: consts.CreditCardModelName, Controller: controllers.NewCreditCardController()},
		{ModelName: consts.BankAccountModelName, Controller: controllers.NewBankAccountControler()},
		{ModelName: consts.LanguageModelName, Controller: controllers.NewLanguageController()},
	}
	return modelList
}
