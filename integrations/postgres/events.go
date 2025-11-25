package postgres

import (
	"database/sql"
	"log"
	"strconv"

	"stintmaster/api/domains/events/normalizers"
)

func CreateEvent(conn *sql.DB, event normalizers.Event) (id int64, err error) {

	query := `INSERT INTO event (name, platform, date, duration, image_url, created_by) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	err = conn.QueryRow(query, event.Name, event.Platform, event.Date, event.Duration, event.ImageURL, event.CreatedBy).Scan(&id)
	if err != nil {
		log.Println("Error executing query:", err)
		return 0, err
	}

	return id, nil
}

func GetEvents(conn *sql.DB) (events []normalizers.Event, err error) {

	query := `SELECT id, name, platform, date, duration, image_url, created_by, created_at FROM event`

	rows, err := conn.Query(query)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var event normalizers.Event
		err := rows.Scan(&event.ID, &event.Name, &event.Platform, &event.Date, &event.Duration, &event.ImageURL, &event.CreatedBy, &event.CreatedAt)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		log.Println("Row iteration error:", err)
		return nil, err
	}

	return events, nil
}

func GetEventsByFilter(conn *sql.DB, filter normalizers.Event) (events []normalizers.Event, err error) {

	query := `SELECT id, name, platform, date, duration, image_url, created_by, created_at FROM event WHERE 1=1`
	var args []interface{}
	argID := 1

	if filter.Name != "" {
		query += ` And (name ILIKE $` + strconv.Itoa(argID)
		args = append(args, "%"+filter.Name+"%")
		argID++
	}
	if filter.Platform != "" {
		query += ` And platform ILIKE $` + strconv.Itoa(argID)
		args = append(args, "%"+filter.Platform+"%")
		argID++
	}
	if !filter.Date.IsZero() {
		query += ` And date = $` + strconv.Itoa(argID)
		args = append(args, filter.Date)
		argID++
	}
	if filter.Duration != 0 {
		query += ` And duration = $` + strconv.Itoa(argID) + `)`
		args = append(args, filter.Duration)
		argID++
	}
	if filter.CreatedBy != "" {
		query += ` or created_by ILIKE $` + strconv.Itoa(argID)
		args = append(args, "%"+filter.CreatedBy+"%")
		argID++
	}

	rows, err := conn.Query(query, args...)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var event normalizers.Event
		err := rows.Scan(&event.ID, &event.Name, &event.Platform, &event.Date, &event.Duration, &event.ImageURL, &event.CreatedBy, &event.CreatedAt)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		log.Println("Row iteration error:", err)
		return nil, err
	}

	return events, nil
}
