package pilots

import (
	"log"
	apin "stintmaster/api/api/v1/pilots/normalizers"
	"stintmaster/api/domains/pilots/normalizers"
	"stintmaster/api/integrations/postgres"
)

func CreatePilot(pilot apin.PostPilot) (id int64, err error) {

	conn := postgres.OpenConnection()

	defer conn.Close()

	reqPilot := normalizers.Pilot{
		Name:       pilot.Name,
		Age:        pilot.Age,
		Experience: pilot.Experience,
		Team:       pilot.Team,
		IracingID:  pilot.IracingID,
		CreatedBy:  pilot.CreatedBy,
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
