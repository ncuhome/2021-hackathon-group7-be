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

func CreateActivity(req *dto.Activities) uint {
	len := utf8.RuneCountInString(req.Content)
	minn := math.Min(60, float64(len))
	dig := string([]rune(req.Content)[:int(minn)])
	activity := dao.Activity{
		Title:     req.Title,
		Content:   req.Content,
		UserId:    req.UserId,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Place:     req.Place,
		Digest:    dig,
	}
	err := activity.Create()
	if err != nil {
		return 1
	}
	_ = activity.SetCacheActivity()
	err = activity.DelCacheList("list")
	if err != nil {
		return 1
	}
	err = activity.DelCacheList("place")
	if err != nil {
		return 1
	}
	err = activity.DelCacheList("host")
	if err != nil {
		return 1
	}
	return 0
}

func GetAllActivities() (interface{}, uint) {
	activity := &dao.Activity{}
	var data []interface{}
	var datareturn interface{}
	list, err := activity.GetCacheAllActivities()
	if err == nil {
		for _, t := range list {
			temp, err := strconv.Atoi(t)
			activity.ID = uint(temp)
			s, err := activity.GetCacheActivity()

			if err != nil {
				p, err := activity.GetActivity()
				p.SetCacheActivity()
				if err != nil {
					return nil, 1
				}
				data = append(data, p)
			} else {
				data = append(data, s)
			}

		}
		datareturn = data

	} else {
		err = activity.SetCacheActivityList()
		datareturn, err = activity.GetAllActivities()
		if err != nil {
			fmt.Println(err)
			return nil, 1
		}
	}
	return datareturn, 0
}

func GetActivity(id string) (interface{}, uint) {

	temp, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println(err)
		return dao.Activity{}, 1
	}

	ID := uint(temp)
	activity := &dao.Activity{
		Model: gorm.Model{
			ID: ID,
		},
	}

	data, err := activity.GetCacheActivity()
	if err == nil {
		return data, 0
	}

	data1, err := activity.GetActivity()
	data1.SetCacheActivity()
	if err != nil {
		fmt.Println(err)
		return data1, 1
	}
	return data1, 0
}

func GetActivitiesByPlace(place string) (interface{}, uint) {
	activity := &dao.Activity{
		Place: place,
	}
	var data []interface{}
	var datareturn interface{}
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
					return nil, 1
				}
				data = append(data, p)
			} else {
				data = append(data, s)
			}

		}
		datareturn = data

	} else {
		err = activity.SetCachePlaceList()
		datareturn, err = activity.GetActivitiesByPlace()
		if err != nil {
			fmt.Println(err)
			return nil, 1
		}
	}
	return datareturn, 0
}

func GetActivitiesByHost(host string) (interface{}, uint) {
	activity := &dao.Activity{
		UserId: host,
	}
	var data []interface{}
	var datareturn interface{}
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
					return nil, 1
				}
				data = append(data, p)
			} else {
				data = append(data, s)
			}

		}
		datareturn = data

	} else {
		err = activity.SetCacheHostList()
		datareturn, err = activity.GetActivitiesByHost()
		if err != nil {
			fmt.Println(err)
			return nil, 1
		}
	}
	return datareturn, 0
}
