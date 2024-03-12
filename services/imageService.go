package services

import (
	"image"
	"image/jpeg"
	"mime/multipart"
	"os"

	compression "github.com/nurlantulemisov/imagecompression"
)

type IImageService interface {
	StoreImage(image image.Image, destination string) error
	CompressImage(image image.Image) image.Image
}

type IImageThirdPartyService interface {
	UploadImage(image *multipart.File)
}

type ImageService struct {
	compressionLevel int
}

func (imageService *ImageService) CompressImage(image image.Image) image.Image {
	compressing, _ := compression.New(imageService.compressionLevel)
	compressingImage := compressing.Compress(image)
	return compressingImage
}

func (imageService *ImageService) StoreImage(image image.Image, destination string) error {
	f, err := os.Create(destination)
	if err != nil {
		return err
	}

	err = jpeg.Encode(f, image, &jpeg.Options{Quality: 60})
	if err != nil {
		return err
	}
	return f.Close()
}

// ##DEPRECATED
func NewImageService(compressionLevel int) *ImageService {
	return &ImageService{
		compressionLevel: compressionLevel,
	}
}
