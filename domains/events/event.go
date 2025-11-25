package events

import (
	"errors"
	"log"
	"time"

	apin "stintmaster/api/api/v1/events/normalizers"
	"stintmaster/api/domains/events/normalizers"
	"stintmaster/api/integrations/postgres"
)

func CreateEvent(event apin.PostEvent) (id int64, err error) {

	conn := postgres.OpenConnection()

	defer conn.Close()

	reqEvent, err := normalizers.EventFromPostEvent(event, time.Now())
	if err != nil {
		log.Println("Error normalizing event:", err)
		return 0, err
	}

	events, err := GetEventsByFilter(event)

	if err != nil {
		log.Println("Error checking existing events:", err)
		return 0, err
	}

	for _, e := range events {
		if e.Name == reqEvent.Name && e.Date.Equal(reqEvent.Date) && e.Platform == reqEvent.Platform {
			log.Println("Event already exists with the same name, date, and platform")
			return 0, errors.New("event already exists with the same name, date, and platform")
		}
	}

	// Limit the number of events and prevent spamming
	if len(events) > 5 {
		log.Println("Event limit exceeded for the given filter")
		return 0, errors.New("event limit exceeded for the given filter")
	}

	id, err = postgres.CreateEvent(conn, reqEvent)

	if err != nil {
		log.Println("Error creating event in database:", err)
		return 0, err
	}

	return id, nil
}

func GetEvents() (events []normalizers.Event, err error) {

	conn := postgres.OpenConnection()

	defer conn.Close()

	events, err = postgres.GetEvents(conn)

	if err != nil {
		log.Println("Error getting events from database:", err)
		return nil, err
	}

	return events, nil
}

func GetEventsByFilter(event apin.PostEvent) (events []normalizers.Event, err error) {

	conn := postgres.OpenConnection()

	defer conn.Close()

	reqEvent, err := normalizers.EventFromPostEvent(event, time.Now())
	if err != nil {
		log.Println("Error normalizing event:", err)
		return nil, err
	}

	events, err = postgres.GetEventsByFilter(conn, reqEvent)

	if err != nil {
		log.Println("Error getting events from database by filter:", err)
		return nil, err
	}

	return events, nil
}
