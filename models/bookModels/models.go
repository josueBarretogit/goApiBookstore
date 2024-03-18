package bookmodels

import (

	"gorm.io/gorm"
)


type DigitalFormat struct {
	gorm.Model
	Price        float64 `json:"price,omitempty"`
	ScreenReader bool    `json:"screenReader,omitempty"`
	TextToSpeech bool    `json:"textToSpeech,omitempty"`
	BookID       *uint   `json:"book_id,omitempty"`
}

type HardCoverFormat struct {
	gorm.Model
	Height   float32 `json:"height,omitempty"`
	Width    float32 `json:"width,omitempty"`
	Length   float32 `json:"length,omitempty"`
	Price    float64 `json:"price,omitempty"`
	NumPages int     `json:"numPages,omitempty"`
	Stock    int64   `json:"stock,omitempty"`
	Weight   float64 `json:"weight,omitempty"`
	BookID   *uint   `json:"idBook,omitempty"`
}

type AudioBookFormat struct {
	gorm.Model
	Price       float64 `json:"price,omitempty"`
	Duration    string  `json:"duration,omitempty"`
	ProgramType string  `json:"programType,omitempty"`
	BookID      *uint   `json:"idBook,omitempty"`
}




