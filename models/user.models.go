package models

import (
	"time"

	"gorm.io/gorm"
)


type User struct {
	ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt 
	Name string  `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	ConfrimPassword string `json:"confrim_password"`
}

