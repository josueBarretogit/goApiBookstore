package bookmodels

import (
	paymentmodels "api/bookstoreApi/models/paymentModels"
	usermodels "api/bookstoreApi/models/userModels"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title           string                          `json:"title,omitempty"`
	CoverPhotoUrl   string                          `json:"cover_photo_url,omitempty"`
	Description     string                          `json:"description,omitempty"`
	Rating          int                             `json:"rating,omitempty"`
	PublicationDate time.Time                       `json:"publication_date,omitempty"`
	Genre           string                          `json:"genre,omitempty"`
	Language        string                          `json:"language,omitempty"`
	ISBN            string                          `json:"isbn,omitempty"`
	Ranking         string                          `json:"ranking,omitempty"`
	AuthorID        *uint                           `json:"author_id,omitempty"`
	Author          usermodels.Author               `json:"author,omitempty"`
	PurchaseDetails []paymentmodels.PurchaseDetails `json:"purchase_details,omitempty"`
}

type DigitalFormat struct {
	gorm.Model
	Price        float64 `json:"price,omitempty"`
	ScreenReader bool    `json:"screen_reader,omitempty"`
	TextToSpeech bool    `json:"text_to_speech,omitempty"`
	BookFormatID *uint   `json:"book_format_id,omitempty"`
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
