package normalizers

import "time"

type PostPilot struct {
	Name         string        `json:"name" validate:"required"`
	Irating      int           `json:"irating"`
	Cars         []string      `json:"cars"`
	Restrictions []Restriction `json:"restrictions"`
}

type Restriction struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}
