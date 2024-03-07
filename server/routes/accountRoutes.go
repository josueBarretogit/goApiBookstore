package routes

import (
	"api/bookstoreApi/controllers"

	"github.com/gin-gonic/gin"
)


func SetupRoutesAccount(r *gin.Engine)  {

	accountrRoutes := r.Group("account")
	accountController := controllers.NewAccountController()
	{
		accountrRoutes.POST("/logIn", accountController.LogIn())
		accountrRoutes.POST("/register", accountController.Register())
		accountrRoutes.PUT("/assignRole/:id",  accountController.AssignRole())
	}

}
