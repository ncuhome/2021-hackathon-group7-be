package service

import (
	"fmt"
	"gorm.io/gorm"
	"math"
	"strconv"
	"tudo/model/dao"
	"tudo/model/dto"
	"unicode/utf8"
)

func CheckV(id uint) uint {
	userInfo := &dao.UserInfo{
		UserID: id,
	}
	err := userInfo.Retrieve()
	if err != nil {
		return ErrorServer
	}
	if userInfo.Verification != "v" {
		return ErrorToken
	}
	return SuccessCode
}

func CreateActivity(req *dto.Activities, id uint) uint {
	code := CheckV(id)
	if code != SuccessCode {
		return code
	}
	contentLen := utf8.RuneCountInString(req.Content)
	minn := math.Min(60, float64(contentLen))
	dig := string([]rune(req.Content)[:int(minn)])
	activity := dao.Activity{
		Title:     req.Title,
		Content:   req.Content,
		UserId:    id,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Place:     req.Place,
		Digest:    dig,
	}
	err := activity.Create()
	if err != nil {
		return ErrorServer
	}
	_ = activity.SetCacheActivity()
	err = activity.DelCacheList("list")
	if err != nil {
		return ErrorServer
	}
	err = activity.DelCacheList("place")
	if err != nil {
		return ErrorServer
	}
	err = activity.DelCacheList("host")
	if err != nil {
		return ErrorServer
	}
	return SuccessCode
}

func GetAllActivities() (interface{}, uint) {
	activity := &dao.Activity{}
	var data []interface{}
	var DataReturn interface{}
	list, err := activity.GetCacheAllActivities()
	if err == nil {
		for _, t := range list {
			temp, err := strconv.Atoi(t)
			activity.ID = uint(temp)
			s, err := activity.GetCacheActivity()

			if err != nil {
				p, err := activity.GetActivity()
				_ = p.SetCacheActivity()
				if err != nil {
					return nil, ErrorServer
				}
				data = append(data, p)
			} else {
				data = append(data, s)
			}

		}
		DataReturn = data

	} else {
		err = activity.SetCacheActivityList()
		DataReturn, err = activity.GetAllActivities()
		if err != nil {
			fmt.Println(err)
			return nil, ErrorServer
		}
	}
	return DataReturn, SuccessCode
}

func GetActivity(id string) (interface{}, uint) {

	temp, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println(err)
		return dao.Activity{}, ErrorServer
	}

	ID := uint(temp)
	activity := &dao.Activity{
		Model: gorm.Model{
			ID: ID,
		},
	}

	data, err := activity.GetCacheActivity()
	if err == nil {
		return data, SuccessCode
	}

	data1, err := activity.GetActivity()
	data1.SetCacheActivity()
	if err != nil {
		fmt.Println(err)
		return data1, ErrorServer
	}
	return data1, SuccessCode
}

func GetActivitiesByPlace(place string) (interface{}, uint) {
	activity := &dao.Activity{
		Place: place,
	}
	var data []interface{}
	var DataReturn interface{}
	list, err := activity.GetCacheActivitiesByPlace()
	if err == nil {
		for _, t := range list {
			temp, err := strconv.Atoi(t)
			activity.ID = uint(temp)
			s, err := activity.GetCacheActivity()

			if err != nil {
				p, err := activity.GetActivity()
				p.SetCacheActivity()
				if err != nil {
					return nil, ErrorServer
				}
				data = append(data, p)
			} else {
				data = append(data, s)
			}

		}
		DataReturn = data

	} else {
		err = activity.SetCachePlaceList()
		DataReturn, err = activity.GetActivitiesByPlace()
		if err != nil {
			fmt.Println(err)
			return nil, ErrorServer
		}
	}
	return DataReturn, SuccessCode
}

func GetActivitiesByHost(host uint) (interface{}, uint) {
	activity := &dao.Activity{
		UserId: host,
	}
	var data []interface{}
	var DataReturn interface{}
	list, err := activity.GetCacheActivitiesByHost()
	if err == nil {
		for _, t := range list {
			temp, err := strconv.Atoi(t)
			activity.ID = uint(temp)
			s, err := activity.GetCacheActivity()

			if err != nil {
				p, err := activity.GetActivity()
				p.SetCacheActivity()
				if err != nil {
					return nil, ErrorServer
				}
				data = append(data, p)
			} else {
				data = append(data, s)
			}

		}
		DataReturn = data

	} else {
		err = activity.SetCacheHostList()
		DataReturn, err = activity.GetActivitiesByHost()
		if err != nil {
			fmt.Println(err)
			return nil, ErrorServer
		}
	}
	return DataReturn, SuccessCode
}
