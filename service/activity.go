package service

import (
	"fmt"
	"gorm.io/gorm"
	"math"
	"nspyf/model/dao"
	"nspyf/model/dto"
	"strconv"
	"unicode/utf8"
)

func CheckV(id uint) uint {
	userInfo := &dao.UserInfo{
		UserID: id,
	}
	err := userInfo.Retrieve()
	if err != nil {
		return ServerError
	}
	if userInfo.Verification != "v" {
		return TokenError
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
		return ServerError
	}
	fmt.Println(666)
	_ = activity.SetCacheActivity()
	err = activity.DelCacheList("list")
	if err != nil {
		return ServerError
	}
	err = activity.DelCacheList("place")
	if err != nil {
		return ServerError
	}
	err = activity.DelCacheList("host")
	if err != nil {
		return ServerError
	}
	return SuccessCode
}

func GetAllActivities(pre int) (interface{}, uint) {
	activity := &dao.Activity{}
	var data []interface{}
	var DataReturn interface{}
	list, err := activity.GetCacheAllActivities(pre)
	if err == nil {
		for i := len(list) - 1; i >= 0; i-- {
			t := list[i]
			temp, err := strconv.Atoi(t)
			activity.ID = uint(temp)
			s, err := activity.GetCacheActivity()

			if err != nil {
				p, err := activity.GetActivity()
				_ = p.SetCacheActivity()
				if err != nil {
					return nil, ServerError
				}
				data = append(data, p)
			} else {
				data = append(data, s)
			}

		}
		DataReturn = data

	} else {
		err = activity.SetCacheActivityList(pre)
		DataReturn, err = activity.GetAllActivities(pre)
		if err != nil {
			fmt.Println(err)
			return nil, ServerError
		}
	}
	return DataReturn, SuccessCode
}

func GetActivity(id string) (interface{}, uint) {

	temp, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println(err)
		return dao.Activity{}, ServerError
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
	_ = data1.SetCacheActivity()
	if err != nil {
		fmt.Println(err)
		return data1, ServerError
	}
	return data1, SuccessCode
}

func GetActivitiesByPlace(place string, pre int) (interface{}, uint) {
	activity := &dao.Activity{
		Place: place,
	}
	var data []interface{}
	var DataReturn interface{}
	list, err := activity.GetCacheActivitiesByPlace(pre)
	if err == nil {
		for i := len(list) - 1; i >= 0; i-- {
			t := list[i]
			temp, err := strconv.Atoi(t)
			activity.ID = uint(temp)
			s, err := activity.GetCacheActivity()

			if err != nil {
				p, err := activity.GetActivity()
				_ = p.SetCacheActivity()
				if err != nil {
					return nil, ServerError
				}
				data = append(data, p)
			} else {
				data = append(data, s)
			}

		}
		DataReturn = data

	} else {
		err = activity.SetCachePlaceList(pre)
		DataReturn, err = activity.GetActivitiesByPlace(pre)
		if err != nil {
			fmt.Println(err)
			return nil, ServerError
		}
	}
	return DataReturn, SuccessCode
}

func GetActivitiesByHost(host uint, pre int) (interface{}, uint) {
	activity := &dao.Activity{
		UserId: host,
	}
	var data []interface{}
	var DataReturn interface{}
	list, err := activity.GetCacheActivitiesByHost(pre)
	if err == nil {
		for i := len(list) - 1; i >= 0; i-- {
			t := list[i]
			temp, err := strconv.Atoi(t)
			activity.ID = uint(temp)
			s, err := activity.GetCacheActivity()

			if err != nil {
				p, err := activity.GetActivity()
				_ = p.SetCacheActivity()
				if err != nil {
					return nil, ServerError
				}
				data = append(data, p)
			} else {
				data = append(data, s)
			}

		}
		DataReturn = data

	} else {
		err = activity.SetCacheHostList(pre)
		DataReturn, err = activity.GetActivitiesByHost(pre)
		if err != nil {
			fmt.Println(err)
			return nil, ServerError
		}
	}
	return DataReturn, SuccessCode
}
