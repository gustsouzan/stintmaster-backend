package postgres

import (
	"log"
	"stintmaster/api/integrations/postgres/models"
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
