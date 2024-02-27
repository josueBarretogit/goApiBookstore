package bookmodels

import (
	paymentmodels "api/bookstoreApi/models/paymentModels"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title           string                          `json:"title,omitempty"`
	CoverUrl        string                          `json:"cover_url,omitempty"`
	Description     string                          `json:"description,omitempty"`
	numPages        int                             `json:"num_pages,omitempty"`
	Rating          int                             `json:"rating,omitempty"`
	PublicationDate time.Time                       `json:"publication_date,omitempty"`
	Genre           string                          `json:"genre,omitempty"`
	Language        string                          `json:"language,omitempty"`
	ISBN            string                          `json:"isbn,omitempty"`
	Ranking         string                          `json:"ranking,omitempty"`
	Stock           int64                           `json:"stock,omitempty"`
	AuthorID        uint                            `json:"author_id,omitempty"`
	BookFormat      []BookFormat                    `json:"book_format,omitempty"`
	PurchaseDetails []paymentmodels.PurchaseDetails `json:"purchase_details,omitempty"`
}

type BookFormat struct {
	gorm.Model
	FormatName      string          `json:"format_name,omitempty"`
	Price           float64         `json:"price,omitempty"`
	DigitalFormat   DigitalFormat   `json:"digital_format,omitempty"`
	HardCoverFormat HardCoverFormat `json:"hard_cover_format,omitempty"`
	BookID          uint            `json:"book_id,omitempty"`
}

type DigitalFormat struct {
	gorm.Model
	FormatName   string  `json:"format_name,omitempty"`
	Price        float64 `json:"price,omitempty"`
	BookFormatID uint    `json:"book_format_id,omitempty"`
}

type HardCoverFormat struct {
	gorm.Model
	Height       float32 `json:"height,omitempty"`
	Width        float32 `json:"width,omitempty"`
	Length       float32 `json:"length,omitempty"`
	BookFormatID uint    `json:"book_format_id,omitempty"`
}
