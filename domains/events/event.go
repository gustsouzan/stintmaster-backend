package events

import (
	"log"
	"strconv"
	"time"

	apin "stintmaster/api/api/v1/events/normalizers"
	"stintmaster/api/domains/events/normalizers"
	"stintmaster/api/integrations/postgres"
)

func CreateEvent(event apin.PostEvent) (id int64, err error) {

	conn := postgres.OpenConnection()

	defer conn.Close()

	loc := time.FixedZone("UTC-3", -3*60*60)
	layout := "2006-01-02"
	t, err := time.ParseInLocation(layout, event.Date, loc)
	if err != nil {
		log.Println("Error parsing date:", err)
		return 0, err
	}

	duration, err := strconv.Atoi(event.Duration)
	if err != nil {
		log.Println("Error parsing duration:", err)
		return 0, err
	}

	reqEvent := normalizers.Event{
		Name:      event.Name,
		Platform:  event.Platform,
		Date:      t,
		Duration:  duration,
		CreatedBy: event.CreatedBy,
		ImageURL:  event.ImageURL,
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
