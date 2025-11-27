package postgres

import (
	"stintmaster/api/integrations/postgres/models"

	"github.com/gofiber/fiber/v2/log"
)

type PilotRepository struct{}

func NewPilotRepository() *PilotRepository {
	return &PilotRepository{}
}

func (r *PilotRepository) CreatePilot(pilot *models.Piloto) error {

	result := dbInstance.Create(&pilot)
	if result.Error != nil {
		log.Error("Erro ao criar piloto:", result.Error)
		return result.Error
	}

	return nil
}

func (r *PilotRepository) GetPilots() (pilots []models.Piloto, err error) {

	result := dbInstance.Find(&pilots)
	if result.Error != nil {
		log.Error("Erro ao consultar pilotos:", result.Error)
		return nil, result.Error
	}

	return pilots, nil
}

func (r *PilotRepository) GetPilotsByFilter(filter *models.Piloto) (pilots []models.Piloto, err error) {

	db := dbInstance

	if filter.IracingID != "" {
		db = db.Where("iracing_id = ?", filter.IracingID)
	}

	if filter.Email != "" {
		db = db.Or("email = ?", filter.Email)
	}

	result := db.Find(&pilots)
	if result.Error != nil {
		log.Error("Erro ao consultar por filtros:", result.Error)
		return nil, result.Error
	}

	return pilots, nil
}
