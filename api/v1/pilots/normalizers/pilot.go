package normalizers

type PostPilot struct {
	Name       string `json:"name" validate:"required"`
	Age        int    `json:"age" validate:"required,gte=0"`
	Experience int    `json:"experience" validate:"required"`
	Team       string `json:"team" validate:"required"`
	IracingID  string `json:"iracing_id" validate:"required"`
	CreatedBy  string `json:"created_by" validate:"required"`
}
