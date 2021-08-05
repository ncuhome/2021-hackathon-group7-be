package dto

type Activity struct {
	Title     string `json:"title" form:"title" binding:"required"`
	Content   string `json:"content" form:"content" binding:"required"`
	// Time：自1970年1月1日00:00:00 UTC以来经过的毫秒数
	StartTime string `json:"start_time" form:"start_time" binding:"required"`
	EndTime   string `json:"end_time" form:"end_time" binding:"required"`
	Place     string `json:"place" form:"place" binding:"required"`
}

type Entity struct {
	ID     string `json:"id" form:"id" binding:"required"`
}
