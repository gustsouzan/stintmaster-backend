package normalizers

import (
	"stintmaster/api/integrations/postgres/models"
	"time"
)

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

type CreatedPilots struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Irating int    `json:"irating"`
}

func PilotDomainToCreatedPilots(pilotDomain []models.Piloto) []CreatedPilots {
	var createdPilots []CreatedPilots
	for _, p := range pilotDomain {
		createdPilots = append(createdPilots, CreatedPilots{
			ID:      p.ID,
			Name:    p.Nome,
			Irating: p.Irating,
		})
	}
	return createdPilots

}
