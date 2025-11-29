package track

import (
	"log"
	"stintmaster/api/integrations/postgres"
	"stintmaster/api/integrations/postgres/models"
)

func GetTracks() (tracks []models.Pista, err error) {

	repository := postgres.NewTrackRepository()
	tracks, err = repository.GetAllTracks()
	if err != nil {
		log.Println("Error retrieving tracks from database:", err)
		return []models.Pista{}, err
	}
	return tracks, err
}
