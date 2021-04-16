package dto

type Activities struct {
	Title     string `json:"title" form:"title" binding:"required"`
	Content   string `json:"content" form:"content" binding:"required"`
	StartTime string `json:"start_time" form:"start_time" binding:"required"`
	EndTime   string `json:"end_time" form:"end_time" binding:"required"`
	Place     string `json:"place" form:"place" binding:"required"`
}
