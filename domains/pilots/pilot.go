package pilots

import (
	"errors"
	"log"
	apin "stintmaster/api/api/v1/pilots/normalizers"
	"stintmaster/api/domains/pilots/normalizers"
	"stintmaster/api/integrations/postgres"
	"stintmaster/api/integrations/postgres/models"
)

func CreatePilot(pilot apin.PostPilot) (id int64, err error) {

	reqPilot, err := normalizers.PostPilotToModal(pilot)
	if err != nil {
		log.Println("Error normalizing pilot data:", err)
		return 0, err
	}

	pilots, err := GetPilotByFilter(pilot)

	if err != nil {
		log.Println("Error checking existing pilots:", err)
		return 0, err
	}

	if len(pilots) > 0 {
		log.Println("Pilot already exists with the same name or iRacing ID")
		return 0, errors.New("pilot already exists with the same name or iRacing ID")
	}

	repository := postgres.NewPilotRepository()
	err = repository.CreatePilot(&reqPilot)

	if err != nil {
		log.Println("Error creating pilot in database:", err)
		return 0, err
	}

	return id, nil
}

func GetPilots() (pilots []models.Piloto, err error) {

	repository := postgres.NewPilotRepository()
	pilots, err = repository.GetPilots()

	if err != nil {
		log.Println("Error getting pilots from database:", err)
		return nil, err
	}

	return pilots, nil
}

func GetPilotByFilter(pilot apin.PostPilot) (pilots []models.Piloto, err error) {

	reqPilot, err := normalizers.PostPilotToModal(pilot)
	if err != nil {
		log.Println("Error normalizing pilot data:", err)
		return nil, err
	}
	repository := postgres.NewPilotRepository()
	pilots, err = repository.GetPilotsByFilter(&reqPilot)

	if err != nil {
		log.Println("Error getting pilots by filter from database:", err)
		return nil, err
	}

	return pilots, nil
}
