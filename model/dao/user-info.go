package dao

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	UserID			uint	`json:",omitempty" gorm:"index;not null"`
	Nickname		string	`json:",omitempty" gorm:"type:varchar(64);index"`
	Avatar			string	`json:",omitempty" gorm:"type:varchar(255);"`
	Digest			string	`json:",omitempty" gorm:"type:varchar(10000);"`
	Verification	string	`json:",omitempty" gorm:"type:varchar(255);index"`
}

func (s *UserInfo) Create() error {
	return DB.Create(s).Error
}

func (s *UserInfo) Retrieve() error {
	return DB.Model(s).Where(s).First(s).Error
}

func (s *UserInfo) Update(change interface{}) error {
	return DB.Model(s).Where(s).Updates(change).Error
}