package dao

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content			string	`json:",omitempty" gorm:"type:varchar(10000);not null"`
	UserID			uint	`json:",omitempty" gorm:"index;not null"`
	ActivityID		uint	`json:",omitempty" gorm:"index;not null"`
	OrganizationID	uint	`json:",omitempty" gorm:"not null"`
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