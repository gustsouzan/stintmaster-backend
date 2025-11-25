package normalizers

import "time"

type Event struct {
	ID        int64
	Name      string
	Platform  string
	Date      time.Time
	Duration  int64
	CreatedBy string
	CreatedAt time.Time
}
