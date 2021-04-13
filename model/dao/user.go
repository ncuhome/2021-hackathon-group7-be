package dao

import (
	"strconv"
	"time"
)

const (
	UserProfileCacheTime = 10 * time.Minute
)

type User struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	Email       string `json:",omitempty" gorm:"index"`
	Username    string `json:",omitempty" gorm:"type:varchar(16);not null;unique_index"`
	Password    string `json:",omitempty" gorm:"type:varchar(128);not null"`
	Salt        string `json:",omitempty" gorm:"type:varchar(128);not null"`
	LoginStatus string `json:",omitempty" gorm:"type:varchar(16);not null"`
}

type UserProfile struct {
	ID          uint
	LoginStatus string
}

type UserDao struct {
	User    *User
	Profile *UserProfile
}

func (s *User) CreateUser() error {
	return DB.Create(s).Error
}
func (s *User) Create() error {
	userInfo := &UserInfo{
		Nickname:     "",
		Avatar:       "",
		Digest:       "",
		Verification: "",
	}

	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := userInfo.Create(); err != nil {
		tx.Rollback()
		return err
	}
	if err := s.CreateUser(); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (s *User) Retrieve() error {
	return DB.Model(s).Where(s).First(s).Error
}

func (s *User) Update(change interface{}) error {
	return DB.Model(s).Where(s).Updates(change).Error
}

// data传入引用用于接受数据
func (s *User) GetProfile(data interface{}) error {
	return DB.Model(s).Where(s).First(data).Error
}

func (s *UserProfile) SetCache() error {
	cacheObj := &JsonCache{
		Data: s,
		ID:   strconv.Itoa(int(s.ID)),
	}
	return cacheObj.SetData(UserProfileCacheTime)
}

func (s *UserProfile) GetCache(id uint) error {
	cacheObj := &JsonCache{
		Data: s,
		ID:   strconv.Itoa(int(id)),
	}
	return cacheObj.GetData()
}

func (s *UserProfile) DelCache() error {
	cacheObj := &JsonCache{
		Data: s,
		ID:   strconv.Itoa(int(s.ID)),
	}
	return cacheObj.DelData()
}

func (s *UserDao) GetProfile(id uint) error {
	s.Profile = new(UserProfile)
	if s.Profile.GetCache(id) != nil {
		s.User = new(User)
		s.User.ID = id
		err := s.User.GetProfile(s.Profile)
		if err != nil {
			return err
		}
		_ = s.Profile.SetCache()
	}
	return nil
}
