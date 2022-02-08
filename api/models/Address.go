package models

import (
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Address struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Address     string    `gorm:"size:255;not null" json:"address"`
	PostalCode   string    `gorm:"size:255;not null;" json:"postal_code"`
	Latitude  string    `gorm:"not null" json:"latitude"`
	Longitude  string    `gorm:"not null" json:"longitude"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Address) Prepare() {
	p.ID = 0
	p.Address = html.EscapeString(strings.TrimSpace(p.Address))
	p.PostalCode = html.EscapeString(strings.TrimSpace(p.PostalCode))
	p.Latitude = html.EscapeString(strings.TrimSpace(p.Latitude))
	p.PostalCode = html.EscapeString(strings.TrimSpace(p.PostalCode))
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (u *Address) SaveAddress(db *gorm.DB) (*Address, error) {

	var err error = db.Debug().Create(&u).Error
	if err != nil {
		return &Address{}, err
	}
	return u, nil
}
