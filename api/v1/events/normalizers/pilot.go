package normalizers

type PostEvent struct {
	Name      string `json:"name" validate:"required"`
	Platform  string `json:"platform" validate:"required"`
	Date      string `json:"date" validate:"required,datetime=02/01/2006"`
	Duration  int64  `json:"duration" validate:"required"`
	CreatedBy string `json:"created_by" validate:"required"`
}
