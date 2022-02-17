package enrichmentAddress

import (
	"github.com/carloshdurante/geolocation_api/api/models"
	"github.com/carloshdurante/geolocation_api/api/repositories"
	tomtom "github.com/carloshdurante/geolocation_api/api/services/geolocation"
	"gorm.io/gorm"
)

type DbConnection struct {
	Db *gorm.DB
}

func (connection *DbConnection) StartEnrichment(params string, address_id uint64) (*models.Address, error) {
	response, err := tomtom.FetchAddress(params)
	if err != nil {
		return nil, err
	}

	repository := repositories.AddressRepositoryDb{Db: connection.Db}
	address, err := repository.Update(address_id, &models.Address{
		ID:        address_id,
		Latitude:  response.Results[0].Position.Lat,
		Longitude: response.Results[0].Position.Lon,
	})
	if err != nil {
		return nil, err
	}
	return address, nil
}
