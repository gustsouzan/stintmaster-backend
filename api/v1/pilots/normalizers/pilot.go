package normalizers

type PostPilot struct {
	Name       string `json:"name" validate:"required"`
	Email      string `json:"email"`
	Age        string `json:"age" validate:"required,gte=0"`
	Experience string `json:"experience" validate:"required"`
	Team       string `json:"team" validate:"required"`
	IracingID  string `json:"iracing_id" validate:"required"`
	ImageURL   string `json:"image_url"`
	CreatedBy  string `json:"created_by" validate:"required"`
}
