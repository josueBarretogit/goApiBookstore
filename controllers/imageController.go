package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ImageController struct {
	DirectoryToStoreImagesPath string
	Module string
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

	filePath := imageController.DirectoryToStoreImagesPath + imageController.Module + "/" +  id

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

	fileNames := []string{}

	for _, file := range files {
		fileExtension := filepath.Ext(file.Filename)
		filename := strconv.FormatInt(time.Now().UnixMilli(), 10) + fileExtension
		errUpload := c.SaveUploadedFile(file, filePath+ "/" + filename)

		if errUpload != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"response": "There was an error uploading images",
				"error":    errUpload.Error(),
				"success":  false,
				
			})
			return
		}
		fileNames = append(fileNames, filename)
	}
	c.JSON(http.StatusOK, gin.H{
		"response": "Imagenes subidas",
		"pathDirectory": "assets/images/" + imageController.Module + "/" + id ,
		"imagenes": fileNames,
		"success" : true,
	})
}
