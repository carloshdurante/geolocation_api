package models

import (
	"time"
)

type Address struct {
	ID         uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Address    string    `gorm:"size:255;not null" json:"address"`
	PostalCode string    `gorm:"size:255;not null" json:"postal_code"`
	Latitude   float64   `gorm:"not null" json:"latitude"`
	Longitude  float64   `gorm:"not null" json:"longitude"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
