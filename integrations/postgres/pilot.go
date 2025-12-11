package postgres

import (
	"log"
	"stintmaster/api/integrations/postgres/models"

	"gorm.io/gorm/clause"
)

type PilotRepository struct{}

func NewPilotRepository() *PilotRepository {
	return &PilotRepository{}
}

func (r *PilotRepository) CreatePilot(pilot *models.Piloto) error {

	result := dbInstance.Create(&pilot)
	if result.Error != nil {
		log.Print("Error creating pilot:", result.Error)
		return result.Error
	}

	return nil
}

func (r *PilotRepository) GetPilots() ([]models.Piloto, error) {

	var pilot []models.Piloto
	result := dbInstance.Preload("Carros").Find(&pilot)
	if result.Error != nil {
		log.Print("Error fetching pilots:", result.Error)
		return nil, result.Error
	}

	return pilot, nil
}

func (r *PilotRepository) DeletePilot(id uint) error {

	var piloto models.Piloto
	if err := dbInstance.Preload("Carros").Find(&piloto, id).Error; err != nil {
		return err
	}
	if err := dbInstance.Model(&piloto).Association("Carros").Clear(); err != nil {
		return err
	}
	result := dbInstance.Delete(&models.Piloto{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *PilotRepository) GetPilotByCarId(carId uint) ([]models.Piloto, error) {

	var pilots []models.Piloto

	query := dbInstance.Preload(clause.Associations)

	err := query.Joins("Join piloto_carros_disponiveis as a on a.piloto_id = pilotos.id").Where("a.carro_id = ?", carId).Find(&pilots).Error

	if err != nil {
		return nil, err
	}

	return pilots, nil
}
