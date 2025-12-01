package postgres

import (
	"log"
	"os"
	"stintmaster/api/integrations/postgres/models"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func OpenConnection() {

	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PWD")
	dbname := os.Getenv("POSTGRES_DB")

	psqlInfo := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Println("Error opening database connection:", err)
		panic(err)
	}

	err = db.AutoMigrate(&models.Carro{}, &models.Pista{}, &models.Piloto{}, &models.Evento{}, &models.RestricaoHorario{})
	if err != nil {
		log.Println("Error during database migration:", err)
		panic(err)
	}

	dbInstance = db
}
