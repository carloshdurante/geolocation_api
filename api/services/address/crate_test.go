package address_test

import (
	"testing"
	"time"

	"github.com/carloshdurante/geolocation_api/api/models"
	"github.com/carloshdurante/geolocation_api/api/repositories"
	"github.com/carloshdurante/geolocation_api/api/services/address"
)

func TestCreate(t *testing.T)  {
	t.Run("create an adress", func(t *testing.T) {
		want := models.Address{
			ID:         1,
			Address:    "avenida Brasilia",
			PostalCode: "13140000",
			Latitude:   40.0000,
			Longitude:  40.0000,
			CreatedAt:  time.Date(2022, 02, 03, 11,46, 20,10, time.UTC),
			UpdatedAt:  time.Date(2022, 02, 03, 11,46, 20,10, time.UTC),
		}

		repoFake := repositories.FakeRepository{}
		got, _ := address.Create(&want, &repoFake)

		if len(repoFake.Addresses) != 1 {
			t.Errorf("want %v, got %v", 1, len(repoFake.Addresses))
		}

		if got != want {
			t.Errorf("want %v, got %v", want, got) 
		}
	})
}