package controllers

import (
	"api/bookstoreApi/initializers"
	"api/bookstoreApi/models"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	newPost := models.Post{Title: body.Title, Body: body.Title}

	result := initializers.DB.Create(&newPost)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "An error ocurred when creating post",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": newPost,
	})
}

func ReadPost(c *gin.Context) {

	result := initializers.DB.Find(&models.Post)
}
