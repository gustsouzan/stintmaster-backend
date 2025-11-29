package postgres

import "stintmaster/api/integrations/postgres/models"

type TrackRepository struct{}

func NewTrackRepository() *TrackRepository {
	return &TrackRepository{}
}

func (r *TrackRepository) GetAllTracks() (tracks []models.Pista, err error) {
	result := dbInstance.Find(&tracks)
	if result.Error != nil {
		return []models.Pista{}, result.Error
	}
	return tracks, nil
}
