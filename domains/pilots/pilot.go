package pilots

import (
	"log"
	"sort"
	apin "stintmaster/api/api/v1/pilots/normalizers"
	"stintmaster/api/domains/carro"
	"stintmaster/api/domains/pilots/normalizers"
	"stintmaster/api/integrations/postgres"
	"stintmaster/api/integrations/postgres/models"
)

func CreatePilot(pilot apin.PostPilot) (reqPilot models.Piloto, err error) {

	reqPilot, err = normalizers.PostPilotToModal(pilot)
	if err != nil {
		log.Println("Error normalizing pilot data:", err)
		return models.Piloto{}, err
	}

	repository := postgres.NewPilotRepository()
	err = repository.CreatePilot(&reqPilot)

	if err != nil {
		log.Println("Error creating pilot in database:", err)
		return models.Piloto{}, err
	}

	return reqPilot, nil
}

func GetCarSuggestions() ([]normalizers.CarResult, error) {
	pilots, err := GetPilot()
	if err != nil {
		return nil, err
	}

	carCount := make(map[uint]int)
	for _, p := range pilots {
		for _, c := range p.Carros {
			carCount[c.ID]++
		}
	}

	// Filter cars with more than 1 occurrence
	type carStat struct {
		ID  uint
		Cnt int
	}
	var stats []carStat
	for id, cnt := range carCount {
		if cnt > 1 {
			stats = append(stats, carStat{id, cnt})
		}
	}

	// Sort by descending quantity
	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Cnt > stats[j].Cnt
	})

	// Limit to top 3
	if len(stats) > 3 {
		stats = stats[:3]
	}

	var result []normalizers.CarResult
	for _, stat := range stats {
		car, err := carro.GetCarById(stat.ID)
		if err != nil {
			log.Println("Error fetching car by ID:", err)
			continue
		}
		result = append(result, normalizers.CarResult{
			Carro: car,
			Qtd:   stat.Cnt,
		})
	}

	return result, nil
}

func GetPilot() ([]models.Piloto, error) {
	repository := postgres.NewPilotRepository()
	pilots, err := repository.GetPilots()

	if err != nil {
		log.Println("Error fetching pilots from database:", err)
		return nil, err
	}
	return pilots, nil
}
