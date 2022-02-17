package enrichmentAddress

import (
	"log"

	"github.com/carloshdurante/geolocation_api/api/models"
	tomtom "github.com/carloshdurante/geolocation_api/api/services/geolocation"
	"gorm.io/gorm"
)

type apiResponse struct {
	Results []struct {
		Address struct {
			StreetName         string
			ExtendedPostalCode string
		}
		Position struct {
			Lat float64
			Lon float64
		}
	}
}

type DbConnection struct {
	Db *gorm.DB
}

func (connection DbConnection) StartEnrichment(params string, address_id uint) (*models.Address, error) {
	response, err := tomtom.FetchAddress("Padre João Franco, 85")
	if err != nil {
		log.Print("Is not possible enrich address on TOMTOM Api:", err)
	}

	var address models.Address
	if err := connection.Db.First(address, address_id); err != nil {
		return nil, err
	}

	&address.Address = response.Results[0].Address.StreetName
	&address.PostalCode = response.Results[0].Address.ExtendedPostalCode
	&address.Latitude = response.Results[0].Position.Lat
	&address.Longitude = response.Results[0].Position.Lat
	connection.Db.Save(&address)

	return &address, nil
}
