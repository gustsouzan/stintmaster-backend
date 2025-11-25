package normalizers

type PostEvent struct {
	Name      string `json:"name" validate:"required"`
	Platform  string `json:"platform" validate:"required"`
	Date      string `json:"date" validate:"required,datetime=2006-01-02"`
	Duration  string `json:"duration" validate:"required"`
	ImageURL  string `json:"image_url"`
	CreatedBy string `json:"created_by" validate:"required"`
}
