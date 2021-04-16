package dao

import (
	"gorm.io/gorm"
	"strconv"
)

type UserInfo struct {
	gorm.Model
	UserID       uint   `json:",omitempty" gorm:"index;not null"`
	Nickname     string `json:",omitempty" gorm:"type:varchar(64);index"`
	Avatar       string `json:",omitempty" gorm:"type:varchar(255);"`
	Digest       string `json:",omitempty" gorm:"type:varchar(10000);"`
	Verification string `json:",omitempty" gorm:"type:varchar(255);index"`
}

type UserInfoDigest struct {
	UserID       uint   `json:"user_id,omitempty"`
	Nickname     string `json:"nickname,omitempty"`
	Avatar       string `json:"avatar,omitempty"`
	Verification string `json:"verification,omitempty"`
}

type UserInfoDao struct {
	Info *UserInfo
	InfoDigest *UserInfoDigest
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

func (s *UserInfo) GetDigest(data interface{}) error {
	return DB.Model(s).Where(s).First(data).Error
}

func (s *UserInfoDigest) SetCache() error {
	cacheObj := &JsonCache{
		Data: s,
		ID:   strconv.Itoa(int(s.UserID)),
	}
	return cacheObj.SetData(UserDataCacheTime)
}

func (s *UserInfoDigest) GetCache(id uint) error {
	cacheObj := &JsonCache{
		Data: s,
		ID:   strconv.Itoa(int(id)),
	}
	return cacheObj.GetData()
}

func (s *UserInfoDigest) DelCache() error {
	cacheObj := &JsonCache{
		Data: s,
		ID:   strconv.Itoa(int(s.UserID)),
	}
	return cacheObj.DelData()
}

func (s *UserInfoDao) GetDigest(id uint) error {
	s.InfoDigest = new(UserInfoDigest)
	if s.InfoDigest.GetCache(id) != nil {
		s.Info = new(UserInfo)
		s.Info.ID = id
		err := s.Info.GetDigest(s.InfoDigest)
		if err != nil {
			return err
		}
		_ = s.InfoDigest.SetCache()
	}
	return nil
}