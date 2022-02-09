package repositories

import (
	"github.com/carloshdurante/geolocation_api/api/models"
	"gorm.io/gorm"
)

type AddressRepositoryInterface interface {
	Create(address *models.Address) (*models.Address, error)
	FindByID(id uint64) (*models.Address, error)
	FindAll() ([]*models.Address, error)
	Update(id uint64, address *models.Address) (*models.Address, error)
	Delete(id uint64) (*models.Address, error)
}

type AddressRepositoryDb struct {
	Db *gorm.DB
}

func (repo *AddressRepositoryDb) Create(address *models.Address) (*models.Address, error) {
	if err := repo.Db.Create(address).Error; err != nil {
		return nil, err
	}
	return address, nil
}

func (repo *AddressRepositoryDb) FindByID(id uint64) (*models.Address, error) {
	address := &models.Address{}
	if err := repo.Db.First(address, id).Error; err != nil {
		return nil, err
	}
	return address, nil
}

func (repo *AddressRepositoryDb) FindAll() ([]*models.Address, error) {
	var addresses []*models.Address
	if err := repo.Db.Find(&addresses).Error; err != nil {
		return nil, err
	}
	return addresses, nil
}

func (repo *AddressRepositoryDb) Update(id uint64, address *models.Address) (*models.Address, error) {
	if err := repo.Db.Model(address).Where("id = ?", id).Updates(address).Error; err != nil {
		return nil, err
	}
	return address, nil
}

func (repo *AddressRepositoryDb) Delete(id uint64) (*models.Address, error) {
	address := &models.Address{}
	if err := repo.Db.Delete(address, id).Error; err != nil {
		return nil, err
	}
	return address, nil
}
