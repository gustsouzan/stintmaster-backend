package postgres

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func OpenConnection() *sql.DB {

	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PWD")
	dbname := os.Getenv("POSTGRES_DB")

	psqlInfo := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("Error opening database connection:", err)
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		log.Println("Error pinging database:", err)
		panic(err)
	}

	return db
}
