package migrations

import (
	"github.com/carloshdurante/geolocation_api/api/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Address{})
}
