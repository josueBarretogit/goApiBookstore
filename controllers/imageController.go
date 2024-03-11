package controllers

import (
	"api/bookstoreApi/services"
	"bytes"
	"image/jpeg"
	"io"
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
	ImageService services.IImageService
}

func (imageController *ImageController) UploadMultipleImageHandler(c *gin.Context) {
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

		openedFile, errOpen := file.Open()

		if errOpen != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"response": "There was an error opening the file",
				"error":    errOpen.Error(),
				"success":  false,
			})
			return
		}


		fileRead , errFileRead := io.ReadAll(openedFile)

		if errFileRead != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"response": "There was an error reading opened images",
				"error":    errFileRead.Error(),
				"success":  false,
			})
			return
		}

		imageToCompress, errDecoding := jpeg.Decode(bytes.NewReader(fileRead))

		if errDecoding != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"response": "There was an error decoding images",
				"error":    errDecoding.Error(),
				"success":  false,
			})
			return
		}



		errUpload := imageController.ImageService.StoreImage(imageToCompress, filePath+ "/" + filename)


		if errUpload != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"response": "There was an error doing all those things images",
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
	return
}
