package dto

type Email struct {
	Email string `json:"email" form:"email" binding:"required"`
}

type EmailBind struct {
	Email string `json:"email" form:"email" binding:"required"`
	Key   string `json:"key" form:"key" binding:"required"`
}
