package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type ImageController struct {
	DirectoryToStoreImagesPath string
}

func (imageController *ImageController) UploadMultipleImageHadler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"response": "No id found",
			"success":  false,
		})
		return
	}

	filePath := imageController.DirectoryToStoreImagesPath +  id

	errDirectory := os.MkdirAll(filePath, 0755)
	if errDirectory != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"response": "Could not store the images",
			"success":  false,
			"error":    errDirectory.Error(),
		})
		return
	}

	form, _ := c.MultipartForm()
	files := form.File["files"]

	for _, file := range files {
		filename := filepath.Base(file.Filename)
		fmt.Println(filePath+ "/" + filename)
		errUpload := c.SaveUploadedFile(file, filePath+ "/" + filename)

		if errUpload != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"response": "There was an error uploading images",
				"error":    errUpload.Error(),
				"success":  false,
			})
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "Imagenes subidas",
		"pathDirectory": imageController.DirectoryToStoreImagesPath ,
		"imagenes": files,
	})
}
