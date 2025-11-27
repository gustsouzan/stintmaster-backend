package normalizers

import (
	"log"
	apin "stintmaster/api/api/v1/events/normalizers"
	"stintmaster/api/integrations/postgres/models"
	"time"
)

func EventFromPostEvent(postEvent apin.PostEvent) (models.Evento, error) {
	loc := time.FixedZone("UTC-3", -3*60*60)
	layout := "2006-01-02"
	dataEvento, err := time.ParseInLocation(layout, postEvent.Date, loc)
	if err != nil {
		log.Println("Error parsing date:", err)
	}
	event := models.Evento{
		Nome:       postEvent.Name,
		Plataforma: postEvent.Platform,
		DataEvento: dataEvento,
		Duracao:    postEvent.Duration,
		Imagem:     postEvent.Image,
		CreatedBy:  postEvent.CreatedBy,
		PistaId:    postEvent.TrackId,
		Classes:    postEvent.Classes,
		MinPilotos: postEvent.MinDrivers,
		MaxPilotos: postEvent.MaxDrivers,
	}
	return event, nil
}
