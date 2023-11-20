package models

import (
	"gorm.io/gorm"
)

type Genre struct {
	gorm.Model
	Title string `json:"title"`
	Href string `json:"href"`
	Description string `json:"description"`
}
