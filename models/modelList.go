package models

import (

	"gorm.io/gorm"
)




type Prueba struct {
  gorm.Model
	Prueba string `json:"prueba" binding:"required"`
}
