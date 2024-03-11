package bookmodels

import (

	"gorm.io/gorm"
)


type DigitalFormat struct {
	gorm.Model
	Price        float64 `json:"price,omitempty"`
	ScreenReader bool    `json:"screen_reader,omitempty"`
	TextToSpeech bool    `json:"text_to_speech,omitempty"`
	BookID       *uint   `json:"book_id,omitempty"`
}

type HardCoverFormat struct {
	gorm.Model
	Height   float32 `json:"height,omitempty"`
	Width    float32 `json:"width,omitempty"`
	Length   float32 `json:"length,omitempty"`
	Price    float64 `json:"price,omitempty"`
	NumPages int     `json:"num_pages,omitempty"`
	Stock    int64   `json:"stock,omitempty"`
	Weight   float64 `json:"weight,omitempty"`
	BookID   *uint   `json:"book_id,omitempty"`
}

type AudioBookFormat struct {
	gorm.Model
	Price       float64 `json:"price,omitempty"`
	Duration    string  `json:"duration,omitempty"`
	ProgramType string  `json:"program_type,omitempty"`
	BookID      *uint   `json:"book_id,omitempty"`
}




