package controllers

import "github.com/gin-gonic/gin"

type IController interface {
	FindAll(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	FindOneBy(c *gin.Context)
	Delete(c *gin.Context)
}
