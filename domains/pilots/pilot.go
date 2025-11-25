package pilots

import (
	"errors"
	"log"
	apin "stintmaster/api/api/v1/pilots/normalizers"
	"stintmaster/api/domains/pilots/normalizers"
	"stintmaster/api/integrations/postgres"
)

func CreatePilot(pilot apin.PostPilot) (id int64, err error) {

	conn := postgres.OpenConnection()

	defer conn.Close()

	reqPilot, err := normalizers.PostPilotToPilot(pilot)
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

	id, err = postgres.CreatePilot(conn, reqPilot)

	if err != nil {
		log.Println("Error creating pilot in database:", err)
		return 0, err
	}

	return id, nil
}

func GetPilots() (pilots []normalizers.Pilot, err error) {

	conn := postgres.OpenConnection()

	defer conn.Close()

	pilots, err = postgres.GetPilots(conn)

	if err != nil {
		log.Println("Error getting pilots from database:", err)
		return nil, err
	}

	return pilots, nil
}

func GetPilotByFilter(pilot apin.PostPilot) (pilots []normalizers.Pilot, err error) {
	conn := postgres.OpenConnection()

	defer conn.Close()

	reqPilot, err := normalizers.PostPilotToPilot(pilot)
	if err != nil {
		log.Println("Error normalizing pilot data:", err)
		return nil, err
	}
	pilots, err = postgres.GetPilotsByFilter(conn, reqPilot)

	if err != nil {
		log.Println("Error getting pilots by filter from database:", err)
		return nil, err
	}

	return pilots, nil
}
