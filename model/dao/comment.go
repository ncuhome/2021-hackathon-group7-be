package dao

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	gorm.Model
	Content    string `json:",omitempty" gorm:"type:varchar(10000);not null"`
	UserID     uint   `json:",omitempty" gorm:"index;not null"`
	ActivityID uint   `json:",omitempty" gorm:"index;not null"`
}

type CommentRespond struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	Content    string    `json:"content"`
	UserID     uint      `json:"user_id"`
	ActivityID uint      `json:"activity_id"`
}

type CommentList struct {
	Data []CommentRespond
}

func (s *Comment) Create() error {
	return DB.Create(s).Error
}

func (s *Comment) Retrieve() error {
	return DB.Model(s).Where(s).First(s).Error
}

func (s *Comment) Update(change interface{}) error {
	return DB.Model(s).Where(s).Updates(change).Error
}

// 必须传id
func (s *Comment) Delete() error {
	return DB.Model(s).Delete(s).Error
}

func (s *CommentList) RetrieveByActivity(id uint, pre uint) error {
	if pre == 0 {
		return DB.Model(&Comment{}).Where("activity_id = ?", id).Order("id DESC").Limit(10).Find(&(s.Data)).Error
	}
	return DB.Model(&Comment{}).Where("activity_id = ? and id < ?", id, pre).Order("id DESC").Limit(10).Find(&(s.Data)).Error
}

func (s *CommentList) RetrieveByUser(id uint, pre uint) error {
	if pre == 0 {
		return DB.Model(&Comment{}).Where("user_id = ?", id).Order("id DESC").Limit(10).Find(&(s.Data)).Error
	}
	return DB.Model(&Comment{}).Where("user_id = ? and id < ?", id, pre).Order("id DESC").Limit(10).Find(&(s.Data)).Error
}
