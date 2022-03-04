package address

import (
	"github.com/carloshdurante/geolocation_api/api/models"
	"github.com/carloshdurante/geolocation_api/api/repositories"
)

func Create(address *models.Address, r repositories.AddressRepositoryInterface) (models.Address, error) {
	a, err := r.Create(address)
	if err != nil {
		return models.Address{}, err
	}
	return *a, nil
}
