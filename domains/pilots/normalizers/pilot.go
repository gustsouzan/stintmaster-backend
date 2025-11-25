package normalizers

import (
	"log"
	"stintmaster/api/api/v1/pilots/normalizers"
	"strconv"
)

type Pilot struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Experience int    `json:"experience"`
	Team       string `json:"team"`
	IracingID  string `json:"iracing_id"`
	CreatedBy  string `json:"created_by"`
	CreatedAt  string `json:"created_at"`
	ImageURL   string `json:"image_url"`
}

func PostPilotToPilot(pilot normalizers.PostPilot) (Pilot, error) {
	age, err := strconv.Atoi(pilot.Age)
	if err != nil {
		log.Println("Error converting age to integer:", err)
		return Pilot{}, err
	}

	experience, err := strconv.Atoi(pilot.Experience)
	if err != nil {
		log.Println("Error converting experience to integer:", err)
		return Pilot{}, err
	}

	reqPilot := Pilot{
		Name:       pilot.Name,
		Age:        age,
		Experience: experience,
		Team:       pilot.Team,
		IracingID:  pilot.IracingID,
		CreatedBy:  pilot.CreatedBy,
	}
	return reqPilot, nil
}
