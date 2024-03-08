package services

import "mime/multipart"

type IImageService interface {
	StoreImage(image *multipart.File)
}

type IImageThirdPartyService interface {
	UploadImage(image *multipart.File)
}	


