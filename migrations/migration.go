package main

import (
	"api/bookstoreApi/initializers"
	"api/bookstoreApi/models"
)

func init() {
	initializers.LoadEnvVariables()

}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
