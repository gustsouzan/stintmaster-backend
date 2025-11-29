package normalizers

import (
	apin "stintmaster/api/api/v1/events/normalizers"
	"stintmaster/api/integrations/postgres/models"
)

func EventFromPostEvent(postEvent apin.PostEvent) (models.Evento, error) {
	event := models.Evento{
		Duracao: postEvent.Duration,
		Pista: models.Pista{
			ID: postEvent.TrackId,
		},
		Classes:    postEvent.Classes,
		MinPilotos: postEvent.MinDrivers,
		MaxPilotos: postEvent.MaxDrivers,
	}
	return event, nil
}
