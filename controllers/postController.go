package controllers

import (
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
	var Posts []models.Post

	result := initializers.DB.Find(&Posts)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "An error ocurred when reading posts",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": Posts,
	})

}

func UpdatePost(c *gin.Context) {

	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	id := c.Param("id")
	var post models.Post

	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "An error ocurred retriving post to update",
		})
		return
	}

	resultUpdate := initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	if resultUpdate.Error != nil {
		c.JSON(500, gin.H{
			"message": "An error ocurred updating",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": post,
	})

}

func DeletePost(c *gin.Context) {

	var postToDelete models.Post
	id := c.Param("id")

	result := initializers.DB.Delete(&postToDelete, id)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "An error ocurred deleting",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": postToDelete,
	})

}
