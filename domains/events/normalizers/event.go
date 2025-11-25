package normalizers

import (
	"log"
	apin "stintmaster/api/api/v1/events/normalizers"
	"strconv"
	"time"
)

type Event struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Platform  string    `json:"platform"`
	Date      time.Time `json:"date"`
	Duration  int       `json:"duration"`
	ImageURL  string    `json:"image_url"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}

func EventFromPostEvent(postEvent apin.PostEvent, createdAt time.Time) (Event, error) {
	loc := time.FixedZone("UTC-3", -3*60*60)
	layout := "2006-01-02"
	t, err := time.ParseInLocation(layout, postEvent.Date, loc)
	if err != nil {
		log.Println("Error parsing date:", err)
	}

	duration, err := strconv.Atoi(postEvent.Duration)
	if err != nil {
		log.Println("Error parsing duration:", err)
		return Event{}, err
	}

	event := Event{
		Name:      postEvent.Name,
		Platform:  postEvent.Platform,
		Date:      t,
		Duration:  duration,
		ImageURL:  postEvent.ImageURL,
		CreatedBy: postEvent.CreatedBy,
		CreatedAt: createdAt,
	}
	return event, nil
}
