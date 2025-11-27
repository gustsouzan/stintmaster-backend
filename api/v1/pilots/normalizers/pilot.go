package normalizers

type PostPilot struct {
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email"`
	IracingID string `json:"iracing_id" validate:"required"`
	Image     string `json:"image"`
	Irating   string `json:"irating"`
	CreatedBy string `json:"created_by" validate:"required"`
}
