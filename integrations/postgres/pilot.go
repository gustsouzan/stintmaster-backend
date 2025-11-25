package postgres

import (
	"database/sql"
	"stintmaster/api/domains/pilots/normalizers"
	"strconv"
)

func CreatePilot(conn *sql.DB, pilot normalizers.Pilot) (int64, error) {

	query := `INSERT INTO pilot (name, age, experience, team, iracing_id, image_url, created_by)
			  VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	var id int64
	err := conn.QueryRow(query, pilot.Name, pilot.Age, pilot.Experience, pilot.Team, pilot.IracingID, pilot.ImageURL, pilot.CreatedBy).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetPilots(conn *sql.DB) (pilots []normalizers.Pilot, err error) {

	query := `SELECT id, name, age, experience, team, iracing_id, image_url, created_by, created_at FROM pilot`

	rows, err := conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var pilot normalizers.Pilot
		err := rows.Scan(&pilot.ID, &pilot.Name, &pilot.Age, &pilot.Experience, &pilot.Team, &pilot.IracingID, &pilot.ImageURL, &pilot.CreatedBy, &pilot.CreatedAt)
		if err != nil {
			return nil, err
		}
		pilots = append(pilots, pilot)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return pilots, nil
}

func GetPilotsByFilter(conn *sql.DB, filter normalizers.Pilot) (pilots []normalizers.Pilot, err error) {

	query := `SELECT id, name, age, experience, team, iracing_id, created_by, created_at FROM pilot `
	var conditions []string
	var args []interface{}
	if filter.IracingID != "" {
		conditions = append(conditions, `iracing_id = $`+strconv.Itoa(len(args)+1))
		args = append(args, filter.IracingID)
	}

	if len(conditions) > 0 {
		query += "WHERE " + conditions[0]
		for i := 1; i < len(conditions); i++ {
			query += " and " + conditions[i]
		}
	}

	rows, err := conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var pilot normalizers.Pilot
		err := rows.Scan(&pilot.ID, &pilot.Name, &pilot.Age, &pilot.Experience, &pilot.Team, &pilot.IracingID, &pilot.CreatedBy, &pilot.CreatedAt)
		if err != nil {
			return nil, err
		}
		pilots = append(pilots, pilot)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return pilots, nil
}
