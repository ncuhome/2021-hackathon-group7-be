package dto

type PostComment struct {
	Content			string	`json:"content" binding:"required"`
	ActivityID		uint	`json:"activity_id" binding:"required"`
}
