package models

import (
	"gorm.io/gorm"
)

type VinylRecord struct {
	gorm.Model
	ID          uint  `gorm:"primaryKey;autoIncrement"`// Unique identifier for the vinyl record
	Title       string // Title of the vinyl record
	Artist      string // Artist or band name
	ReleaseYear int    // Year of release
	Genre       string // Genre of the music
	Condition   string // Condition of the vinyl record (e.g., "Mint," "Used," etc.)
	Price       float64 // Price of the vinyl record
}