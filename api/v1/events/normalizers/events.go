package normalizers

type PostEvent struct {
	Duration   int      `json:"duration" validate:"required"`
	TrackId    uint     `json:"trackID" validate:"required"`
	Classes    []string `json:"carClasses" validate:"required"`
	MinDrivers int      `json:"minDrivers" validate:"required"`
	MaxDrivers int      `json:"maxDrivers" validate:"required"`
}

type CalculateEventRequest struct {
	PilotID        int     `json:"pilotID" validate:"required"`
	LapTime        string  `json:"lapTime" validate:"required"`
	LapConsumption float64 `json:"lapConsumption" validate:"required"`
	StintMaxLength int     `json:"stintMaxLength" validate:"required"`
}
