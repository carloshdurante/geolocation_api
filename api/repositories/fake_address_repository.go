package repositories

import "github.com/carloshdurante/geolocation_api/api/models"

type FakeRepository struct {
	// Create(address *models.Address) (*models.Address, error)
	// FindByID(id uint64) (*models.Address, error)
	// FindAll() ([]*models.Address, error)
	// Update(id uint64, address *models.Address) (*models.Address, error)
	// Delete(id uint64) (*models.Address, error)
	Addresses []models.Address
}

func (repo *FakeRepository) Create(address *models.Address) (*models.Address, error) {
	// if err := repo.Db.Create(address).Error; err != nil {
	// 	return nil, err
	// }
	// return address, nil
	repo.Addresses = append(repo.Addresses, *address)
	return address, nil
}

func (repo *FakeRepository) FindByID(id uint64) (*models.Address, error) {
	// address := &models.Address{}
	// if err := repo.Db.First(address, id).Error; err != nil {
	// 	return nil, err
	// }
	// return address, nil
	panic("TODO")
}

func (repo *FakeRepository) FindAll() ([]*models.Address, error) {
	// var addresses []*models.Address
	// if err := repo.Db.Find(&addresses).Error; err != nil {
	// 	return nil, err
	// }
	// return addresses, nil
	panic("TODO")
}

func (repo *FakeRepository) Update(id uint64, address *models.Address) (*models.Address, error) {
	// if err := repo.Db.Model(address).Where("id = ?", id).Updates(address).Error; err != nil {
	// 	return nil, err
	// }
	// return address, nil
	panic("TODO")
}

func (repo *FakeRepository) Delete(id uint64) (*models.Address, error) {
	// address := &models.Address{}
	// if err := repo.Db.Delete(address, id).Error; err != nil {
	// 	return nil, err
	// }
	// return address, nil
	panic("TODO")
}
