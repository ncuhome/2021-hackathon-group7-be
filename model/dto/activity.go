package dto

type Activities struct {
	Title     string `json:"title" form:"title" binding:"required"`
	Content   string `json:"content" form:"content" binding:"required"`
	UserId    string `json:"id" form:"id" binding:"required"`
	StartTime string `json:"starttime" form:"starttime" binding:"required"`
	EndTime   string `json:"endtime" form:"endtime" binding:"required"`
	Place     string `json:"place" form:"place" binding:"required"`
}
