package postgres

import (
	"stintmaster/api/integrations/postgres/models"
)

type CarroRepository struct{}

func NewCarroRepository() *CarroRepository {
	return &CarroRepository{}
}

func (r *CarroRepository) GetCarById(id uint) (carro models.Carro, err error) {
	result := dbInstance.First(&carro, id)
	if result.Error != nil {
		return models.Carro{}, result.Error
	}
	return carro, nil
}

func (r *CarroRepository) GetAllCars() (cars []models.Carro, err error) {
	result := dbInstance.Find(&cars)
	if result.Error != nil {
		return []models.Carro{}, result.Error
	}
	return cars, nil
}
