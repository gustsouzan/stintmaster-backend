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

	query := `SELECT id, name, platform, date, duration, image_url, created_by, created_at FROM event`
	var conditions []string
	var args []interface{}

	if filter.Name != "" {
		conditions = append(conditions, "name ILIKE $"+strconv.Itoa(len(args)+1))
		args = append(args, "%"+filter.Name+"%")
	}
	if filter.Platform != "" {
		conditions = append(conditions, "platform ILIKE $"+strconv.Itoa(len(args)+1))
		args = append(args, "%"+filter.Platform+"%")
	}
	if !filter.Date.IsZero() {
		conditions = append(conditions, "date = $"+strconv.Itoa(len(args)+1))
		args = append(args, filter.Date)
	}
	if filter.Duration != 0 {
		conditions = append(conditions, "duration = $"+strconv.Itoa(len(args)+1))
		args = append(args, filter.Duration)
	}
	if filter.CreatedBy != "" {
		conditions = append(conditions, "created_by ILIKE $"+strconv.Itoa(len(args)+1))
		args = append(args, "%"+filter.CreatedBy+"%")
	}

	if len(conditions) > 0 {
		query += " WHERE "
		for i := 0; i < len(conditions); i++ {
			if i > 0 {
				if i == len(conditions)-1 && filter.CreatedBy != "" {
					query += " OR "
				} else {
					query += " AND "
				}
			}
			query += conditions[i]
		}
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
