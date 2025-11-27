package events

import (
	"errors"
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

	events, err := GetEventsByFilter(event)

	if err != nil {
		log.Println("Error checking existing events:", err)
		return models.Evento{}, err
	}

	if len(events) > 5 {
		log.Println("Event limit exceeded for the given filter")
		return models.Evento{}, errors.New("event limit exceeded for the given filter")
	}

	for _, e := range events {
		if e.Nome == reqEvent.Nome && e.DataEvento.Equal(reqEvent.DataEvento) && e.Plataforma == reqEvent.Plataforma {
			log.Println("Event already exists with the same name, date, and platform")
			return models.Evento{}, errors.New("event already exists with the same name, date, and platform")
		}
	}

	repository := postgres.NewEventRepository()
	err = repository.CreateEvent(&reqEvent)

	if err != nil {
		log.Println("Error creating event in database:", err)
		return models.Evento{}, err
	}

	return reqEvent, nil
}

func GetEvents() (events []models.Evento, err error) {

	repository := postgres.NewEventRepository()
	events, err = repository.GetEvents()

	if err != nil {
		log.Println("Error getting events from database:", err)
		return nil, err
	}

	return events, nil
}

func GetEventsByFilter(event apin.PostEvent) (events []models.Evento, err error) {

	reqEvent, err := normalizers.EventFromPostEvent(event)
	if err != nil {
		log.Println("Error normalizing event:", err)
		return nil, err
	}

	repository := postgres.NewEventRepository()
	events, err = repository.GetEventsByFilter(&reqEvent)

	if err != nil {
		log.Println("Error getting events from database by filter:", err)
		return nil, err
	}

	return events, nil
}
