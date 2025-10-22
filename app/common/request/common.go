package request

type ID struct {
	ID uint `form:"id" json:"id" example:"1" binding:"required"`
}
