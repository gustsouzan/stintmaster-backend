package normalizers

import "time"

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
