package bookmodels

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title             string
	CoverUrl          string
	Description       string
	numPages          int
	Rating            int
	PublicationDate   time.Time
	Genre             string
	Language          string
	ISBN              string
	Ranking           string
	Stock             int64
	AuthorID          uint
	PurchaseDetailsID uint
	BookFormat        []BookFormat
}

type BookFormat struct {
	gorm.Model
	FormatName      string
	Price           float64
	DigitalFormat   []DigitalFormat
	HardCoverFormat []HardCoverFormat
	BookID          uint
}

type DigitalFormat struct {
	gorm.Model
	FormatName   string
	Price        float64
	BookFormatID uint
}

type HardCoverFormat struct {
	gorm.Model
	Height       float32
	Width        float32
	Length       float32
	BookFormatID uint
}
