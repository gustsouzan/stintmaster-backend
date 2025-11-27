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

func (r *EventRepository) GetEvents() (events []models.Evento, err error) {

	err = dbInstance.Find(&events).Error
	if err != nil {
		log.Error("Error executing query:", err)
		return nil, err
	}

	return events, nil
}

func (r *EventRepository) GetEventsByFilter(filter *models.Evento) (events []models.Evento, err error) {

	db := dbInstance

	if filter.Nome != "" {
		db = db.Where("nome ILIKE ?", "%"+filter.Nome+"%")
	}
	if filter.Plataforma != "" {
		db = db.Where("plataforma ILIKE ?", "%"+filter.Plataforma+"%")
	}
	if !filter.DataEvento.IsZero() {
		db = db.Where("data_evento = ?", filter.DataEvento)
	}
	if filter.Duracao != 0 {
		db = db.Where("duracao = ?", filter.Duracao)
	}
	if filter.CreatedBy != "" {
		db = db.Or("created_by ILIKE ?", "%"+filter.CreatedBy+"%")
	}

	results := db.Find(&events)
	if results.Error != nil {
		log.Error("Error executing query:", results.Error)
		return nil, results.Error
	}

	return events, nil
}
