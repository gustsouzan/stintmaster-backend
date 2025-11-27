package normalizers

type PostEvent struct {
	Name       string   `json:"name" validate:"required"`
	Platform   string   `json:"platform" validate:"required"`
	Date       string   `json:"date" validate:"required,datetime=2006-01-02"`
	TrackId    uint     `json:"track_id" validate:"required"`
	Duration   int      `json:"duration" validate:"required"`
	Classes    []string `json:"classes" validate:"required"`
	MinDrivers int      `json:"min_drivers" validate:"required"`
	MaxDrivers int      `json:"max_drivers" validate:"required"`
	Image      string   `json:"image"`
	CreatedBy  string   `json:"created_by" validate:"required"`
}
