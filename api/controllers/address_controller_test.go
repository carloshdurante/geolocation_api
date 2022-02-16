package controllers

import (
	"github.com/carloshdurante/geolocation_api/api/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"testing"
)

type AddressRepositoryMock struct {
	mock.Mock
}

func (r *AddressRepositoryMock) GetAllAddress() ([]*models.Address, error) {
	args := r.Called()
	addresses := []*models.Address{
		{ID: 1, Address: "Rua 1", PostalCode: "08031450", Latitude: "-01", Longitude: "01"},
	}
	return addresses, args.Error(1)
}

func TestService_GetAllAddress(t *testing.T) {
	repository := new(AddressRepositoryMock)
	repository.On("GetAllAddress", mock.Anything).Return([]*models.Address{}, nil)

	addresses, _ := repository.GetAllAddress()

	for i := range addresses {
		assert.Equal(t, addresses[i].Address, "Rua 1")
	}

}

func TestService_GetAllAddress_Error(t *testing.T) {
	repository := new(AddressRepositoryMock)
	repository.On("GetAllAddress", mock.Anything).Return([]*models.Address{}, nil)

	_, err := repository.GetAllAddress()

	assert.Nil(t, err)
}

// GET address by id
func TestService_GetAddressByID(t *testing.T) {
	repository := new(AddressRepositoryMock)
	repository.On("GetAllAddress", mock.Anything).Return([]*models.Address{}, nil)

	addresses, _ := repository.GetAllAddress()

	for i := range addresses {
		assert.NotEqual(t, addresses[i].ID, 2, "not be equal")
		assert.Equal(t, addresses[i].ID,  uint32(1), "be equal")
	}
}

// GET error address by id
func TestService_GetAddressByID_Error(t *testing.T) {
	repository := new(AddressRepositoryMock)
	repository.On("GetAllAddress", mock.Anything).Return([]*models.Address{}, nil)

	_, err := repository.GetAllAddress()

	assert.Nil(t, err)
}
