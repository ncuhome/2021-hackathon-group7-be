package dao

import (
	"strconv"
	"time"
)

const (
	UserDataCacheTime = 10 * time.Minute
)

type User struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	Email       string `json:",omitempty" gorm:"index"`
	Phone       string `json:",omitempty" gorm:"index"`
	Username    string `json:",omitempty" gorm:"type:varchar(255);not null;unique_index"`
	Password    string `json:",omitempty" gorm:"type:varchar(128);not null"`
	Salt        string `json:",omitempty" gorm:"type:varchar(128);not null"`
	LoginStatus string `json:",omitempty" gorm:"type:varchar(16);not null"`
}

type UserData struct {
	ID          uint
	LoginStatus string
}

type UserDao struct {
	User *User
	Data *UserData
}

func (s *User) CreateUser() error {
	return DB.Create(s).Error
}
func (s *User) Create() error {
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := s.CreateUser(); err != nil {
		tx.Rollback()
		return err
	}
	userInfo := &UserInfo{
		UserID:       s.ID,
		Nickname:     s.Username,
		Avatar:       "",
		Digest:       "",
		Verification: "",
	}
	if err := userInfo.Create(); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// userInfo 的相关信息写入数据库
func (s *User) CreateWith(userInfo *UserInfo) error { // desc 创建并关联 User.ID 到 UserInfo.UserID
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() // tip 不成功就回滚 保证原子性
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(s).Error; err != nil {
		tx.Rollback()
		return err
	}
	userInfo.UserID = s.ID
	if err := tx.Create(userInfo).Error; err != nil {
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
func (s *User) GetData(data interface{}) error {
	return DB.Model(s).Where(s).First(data).Error
}

func (s *UserData) SetCache() error {
	cacheObj := &JsonCache{
		Data: s,
		ID:   strconv.Itoa(int(s.ID)),
	}
	return cacheObj.SetData(UserDataCacheTime)
}

func (s *UserData) GetCache(id uint) error {
	cacheObj := &JsonCache{
		Data: s,
		ID:   strconv.Itoa(int(id)),
	}
	return cacheObj.GetData()
}

func (s *UserData) DelCache() error {
	cacheObj := &JsonCache{
		Data: s,
		ID:   strconv.Itoa(int(s.ID)),
	}
	return cacheObj.DelData()
}

func (s *UserDao) GetData(id uint) error {
	s.Data = new(UserData)
	if s.Data.GetCache(id) != nil {
		s.User = new(User)
		s.User.ID = id
		err := s.User.GetData(s.Data)
		if err != nil {
			return err
		}
		_ = s.Data.SetCache()
	}
	return nil
}
