package carro

import (
	"stintmaster/api/integrations/postgres"
	"stintmaster/api/integrations/postgres/models"
)

func GetCarById(id uint) (carro models.Carro, err error) {

	repository := postgres.NewCarroRepository()
	carro, err = repository.GetCarById(id)
	if err != nil {
		return models.Carro{}, err
	}
	return carro, nil
}

func GetAllCars() (cars []models.Carro, err error) {

	repository := postgres.NewCarroRepository()
	cars, err = repository.GetAllCars()
	if err != nil {
		return []models.Carro{}, err
	}
	return cars, nil
}

func GetMappedCarClasses() (mappedClasses map[string][]models.Carro, err error) {

	cars, err := GetAllCars()
	if err != nil {
		return nil, err
	}

	mappedClasses = make(map[string][]models.Carro)
	for _, car := range cars {
		mappedClasses[car.Classe] = append(mappedClasses[car.Classe], car)
	}

	return mappedClasses, nil
}

func GetCarClasses() (carClasses []string, err error) {

	cars, err := GetAllCars()
	if err != nil {
		return []string{}, err
	}

	classSet := make(map[string]struct{})
	for _, car := range cars {
		classSet[car.Classe] = struct{}{}
	}

	for class := range classSet {
		carClasses = append(carClasses, class)
	}

	return carClasses, nil
}
