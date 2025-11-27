package normalizers

import (
	"stintmaster/api/api/v1/pilots/normalizers"
	"stintmaster/api/integrations/postgres/models"
	"strconv"
)

func PostPilotToModal(pilot normalizers.PostPilot) (models.Piloto, error) {

	irating, err := strconv.Atoi(pilot.Irating)
	if err != nil {
		return models.Piloto{}, err
	}

	reqPilot := models.Piloto{
		Nome:      pilot.Name,
		Email:     pilot.Email,
		IracingID: pilot.IracingID,
		Imagem:    pilot.Image,
		Irating:   irating,
		CreatedBy: pilot.CreatedBy,
	}
	return reqPilot, nil
}
