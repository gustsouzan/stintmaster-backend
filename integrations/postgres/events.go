package postgres

import (
	"stintmaster/api/integrations/postgres/models"

	"github.com/gofiber/fiber/v2/log"
)

type EventRepository struct{}

func NewEventRepository() *EventRepository {
	return &EventRepository{}
}

func (r *EventRepository) CreateEvent(evento *models.Evento) (err error) {

	result := dbInstance.Create(&evento)
	if result.Error != nil {
		log.Error("Error executing query:", result.Error)
		return result.Error
	}

	return nil
}

func (r *EventRepository) AddPilotToEvent(eventID uint, pilotID uint) error {
	eventPilot := models.EventoPilotosInscritos{
		EventoID: eventID,
		PilotoID: pilotID,
	}

	result := dbInstance.Create(&eventPilot)
	if result.Error != nil {
		log.Error("Error executing query:", result.Error)
		return result.Error
	}

	return nil
}
