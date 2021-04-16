package dao

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

const (
	ActivityCacheTime = 10 * time.Minute
)

type Activity struct {
	gorm.Model
	UserId    uint
	Title     string `gorm:"varchar(128)"`
	Content   string `gorm:"text"`
	StartTime string `gorm:"varchar(64)"`
	EndTime   string `gorm:"varchar(64)"`
	Place     string `gorm:"varchar(128)"`
	Digest    string `gorm:"varchar(255)"`
}

//mysql
//添加新活动
func (s *Activity) Create() error {
	return DB.Create(s).Error
}

//获取所有活动列表
func (s *Activity) GetAllActivities(pre int) ([]Activity, error) {
	var t []Activity
	var err error
	fmt.Println(pre)
	if pre == 0 {
		err = DB.Model(s).Order("id DESC").Limit(10).Find(&t).Error
	} else {
		err = DB.Model(s).Where("id < ?", pre).Order("id DESC").Limit(10).Find(&t).Error
	}

	for _, content := range t {
		_ = content.SetCacheActivity()

	}
	return t, err
}

//获取活动详细信息
func (s *Activity) GetActivity() (Activity, error) {
	var t Activity
	err := DB.Model(s).Find(&t).Error
	return t, err
}

//按地点获取活动
func (s *Activity) GetActivitiesByPlace(pre int) ([]Activity, error) {
	tim := time.Now().Unix()
	var t []Activity
	var err error
	if pre == 0 {
		err = DB.Model(s).Where("place = ?", s.Place).Where("end_time > ?", tim).Order("id DESC").Limit(10).Find(&t).Error
	} else {
		err = DB.Model(s).Where("place = ?", s.Place).Where("id < ?", pre).Where("end_time > ?", tim).Order("id DESC").Limit(10).Find(&t).Error
	}

	for _, content := range t {
		_ = content.SetCacheActivity()

	}
	return t, err
}

//按组织获取活动
func (s *Activity) GetActivitiesByHost(pre int) ([]Activity, error) {
	var t []Activity
	var err error
	if pre == 0 {
		err = DB.Model(s).Where("user_id = ?", s.UserId).Order("id DESC").Limit(10).Find(&t).Error
	} else {
		err = DB.Model(s).Where("user_id = ?", s.UserId).Where("id < ?", pre).Order("id DESC").Limit(10).Find(&t).Error
	}

	for _, content := range t {
		_ = content.SetCacheActivity()

	}
	return t, err
}

//redis
//建立
func (s *Activity) SetCacheActivity() error {
	b := strconv.Itoa(int(s.ID))
	key := CacheConfigObj.Prefix + "activity" + b
	DataBytes, err := json.Marshal(s)
	if err != nil {
		return err
	}
	err = Cache.Set(key, string(DataBytes), ActivityCacheTime).Err()
	Cache.Expire(key, ActivityCacheTime)
	return err
}

func (s *Activity) SetCacheActivityList(pre int) error {
	key := CacheConfigObj.Prefix + "activity_list" + strconv.Itoa(pre)
	data, err := s.GetAllActivities(pre)
	fmt.Println(data)
	if err != nil {
		return err
	}

	for _, temp := range data {
		fmt.Println(temp.ID)
		err = Cache.SAdd(key, temp.ID).Err()
		Cache.Expire(key, ActivityCacheTime)
		if err != nil {
			return err
		}
	}

	return err
}

func (s *Activity) SetCachePlaceList(pre int) error {
	key := CacheConfigObj.Prefix + "activity_place" + strconv.Itoa(pre)
	data, err := s.GetActivitiesByPlace(pre)
	if err != nil {
		return err
	}

	for _, temp := range data {
		err = Cache.SAdd(key, temp.ID).Err()
		Cache.Expire(key, ActivityCacheTime)
		if err != nil {
			return err
		}
	}

	return err
}

func (s *Activity) SetCacheHostList(pre int) error {
	key := CacheConfigObj.Prefix + "activity_host" + strconv.Itoa(pre)
	data, err := s.GetActivitiesByHost(pre)
	if err != nil {
		return err
	}

	for _, temp := range data {
		err = Cache.SAdd(key, temp.ID).Err()
		Cache.Expire(key, ActivityCacheTime)
		if err != nil {
			return err
		}
	}

	return err
}

//获取

func (s *Activity) GetCacheActivity() (interface{}, error) {
	key := CacheConfigObj.Prefix + "activity" + strconv.Itoa(int(s.ID))

	value, err := Cache.Get(key).Result()
	if err != nil {
		return nil, err
	}
	var t interface{}
	err = json.Unmarshal([]byte(value), &t)

	if err != nil {
		return nil, err
	}

	return t, err
}

func (s *Activity) GetCacheAllActivities(pre int) ([]string, error) {
	key := CacheConfigObj.Prefix + "activity_list" + strconv.Itoa(pre)
	members, err := Cache.SCard(key).Result()
	fmt.Println(key, pre)
	if err != nil {
		return nil, err
	}
	if members != 0 {
		val, err := Cache.SMembers(key).Result()
		return val, err
	}

	return nil, errors.New("false")
}

func (s *Activity) GetCacheActivitiesByPlace(pre int) ([]string, error) {
	key := CacheConfigObj.Prefix + "activity_place" + strconv.Itoa(pre)
	members, err := Cache.SCard(key).Result()
	if err != nil {
		return nil, err
	}
	if members != 0 {
		val, err := Cache.SMembers(key).Result()
		return val, err
	}

	return nil, errors.New("false")
}

func (s *Activity) GetCacheActivitiesByHost(pre int) ([]string, error) {
	key := CacheConfigObj.Prefix + "activity_host" + strconv.Itoa(pre)
	members, err := Cache.SCard(key).Result()
	if err != nil {
		return nil, err
	}
	if members != 0 {
		val, err := Cache.SMembers(key).Result()
		return val, err
	}

	return nil, errors.New("false")
}

//删除缓存
func (s *Activity) DelCacheList(name string) error {
	key := CacheConfigObj.Prefix + "activity_" + name + "*"
	var cursor uint64
	var names []string
	for {
		var keys []string
		var err error
		keys, cursor, err = Cache.Scan(cursor, key, 10).Result()
		if err != nil {
			panic(err)
			return err
		}
		for _, p := range keys {
			names = append(names, p)
		}
		if cursor == 0 {
			break
		}
	}
	for _, p := range names {
		err := Cache.Del(p).Err()
		if err != nil {
			return err
		}
	}
	return nil
}
