package events

import (
	"log"

	apin "stintmaster/api/api/v1/events/normalizers"
	"stintmaster/api/domains/events/normalizers"
	"stintmaster/api/integrations/postgres"
	"stintmaster/api/integrations/postgres/models"
)

func CreateEvent(event apin.PostEvent) (models.Evento, error) {

	reqEvent, err := normalizers.EventFromPostEvent(event)
	if err != nil {
		log.Println("Error normalizing event:", err)
		return models.Evento{}, err
	}

	repository := postgres.NewEventRepository()
	err = repository.CreateEvent(&reqEvent)

	if err != nil {
		log.Println("Error creating event in database:", err)
		return models.Evento{}, err
	}

	return reqEvent, nil
}

func JoinPilotToEvent(eventID uint, pilotID uint) error {
	repository := postgres.NewEventRepository()
	err := repository.AddPilotToEvent(eventID, pilotID)

	if err != nil {
		log.Println("Error adding pilot to event in database:", err)
		return err
	}

	return nil
}
