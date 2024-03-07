package routes

import (
	"api/bookstoreApi/consts"
	"api/bookstoreApi/controllers"

	"github.com/gin-gonic/gin"
)


func SetupRoutesAccount(r *gin.Engine)  {

	accountrRoutes := r.Group(consts.AccountModelName)
	accountController := controllers.NewAccountController()
	{
		accountrRoutes.POST("/logIn", accountController.LogIn())
		accountrRoutes.POST("/register", accountController.Register())
		accountrRoutes.PUT("/assignRole/:id",  accountController.AssignRole())
	}

}
