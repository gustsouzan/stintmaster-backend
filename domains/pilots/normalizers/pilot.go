package normalizers

import (
	"stintmaster/api/api/v1/pilots/normalizers"
	"stintmaster/api/integrations/postgres/models"
	"strconv"
)

type CarResult struct {
	models.Carro
	Qtd int
}

func PostPilotToModal(pilot normalizers.PostPilot) (models.Piloto, error) {

	reqPilot := models.Piloto{
		Nome:    pilot.Name,
		Irating: pilot.Irating,
		Carros: func() []models.Carro {
			var cars []models.Carro
			for _, car := range pilot.Cars {
				id, _ := strconv.Atoi(car)
				cars = append(cars, models.Carro{ID: uint(id)})
			}
			return cars
		}(),
		RestricaoHorario: func() []models.RestricaoHorario {
			var restrictions []models.RestricaoHorario
			for _, r := range pilot.Restrictions {
				restrictions = append(restrictions, models.RestricaoHorario{
					HoraInicio: r.StartTime,
					HoraFim:    r.EndTime,
				})
			}
			return restrictions
		}(),
	}
	return reqPilot, nil
}
