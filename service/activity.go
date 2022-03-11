package service

import (
	"math"
	"tudo/model"
	"tudo/model/dao"
	"tudo/model/dto"
	"unicode/utf8"
)

func checkV(id uint) uint {
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

func CreateActivity(req *dto.Activity, id uint) uint {
	code := checkV(id)
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
	return SuccessCode
}

func UpdateActivity(req *dto.Activity, actID uint, userID uint) uint {
	act := &dao.Activity{}
	act.ID = actID
	err := act.Retrieve()
	if err != nil {
		return ErrorCommitData
	}

	if act.UserId != userID {
		return ErrorToken
	}

	contentLen := utf8.RuneCountInString(req.Content)
	minn := math.Min(60, float64(contentLen))
	dig := string([]rune(req.Content)[:int(minn)])
	change := &dao.Activity{
		Title:     req.Title,
		Content:   req.Content,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Place:     req.Place,
		Digest:    dig,
	}

	err = act.Update(change)
	if err != nil {
		model.ErrLog.Println(err)
		return ErrorServer
	}
	return SuccessCode
}

func DeleteActivity(actID uint, userID uint) uint {
	act := &dao.Activity{}
	act.ID = actID
	err := act.Retrieve()
	if err != nil {
		return ErrorCommitData
	}

	if act.UserId != userID {
		return ErrorToken
	}

	err = act.Delete()
	if err != nil {
		model.ErrLog.Println(err)
		return ErrorServer
	}
	return SuccessCode
}

func RetrieveActivity(id uint) (interface{}, uint) {
	act := &dao.ActivityFull{}
	act.ID = id
	err := act.Retrieve()
	if err != nil {
		return nil, ErrorCommitData
	}

	return act, SuccessCode
}

func RetrieveActivityListNotStart(pre int) (interface{}, uint) {
	actList := &dao.ActivityDigestArr{}
	err := actList.RetrieveNotStart(pre, 10)
	if err != nil {
		model.ErrLog.Println(err)
		return nil, ErrorServer
	}

	return actList, SuccessCode
}

func RetrieveActivityListDuring(now int, pre int) (interface{}, uint) {
	actList := &dao.ActivityDigestArr{}
	err := actList.RetrieveDuring(now, pre, 10)
	if err != nil {
		model.ErrLog.Println(err)
		return nil, ErrorServer
	}

	return actList, SuccessCode
}

func RetrieveActivityListEnded(pre int) (interface{}, uint) {
	actList := &dao.ActivityDigestArr{}
	err := actList.RetrieveEnded(pre, 10)
	if err != nil {
		model.ErrLog.Println(err)
		return nil, ErrorServer
	}

	return actList, SuccessCode
}

func RetrieveActivityListEndedByHost(id uint, pre int) (interface{}, uint) {
	actList := &dao.ActivityDigestArr{}
	err := actList.RetrieveEndedByHost(id, pre, 10)
	if err != nil {
		model.ErrLog.Println(err)
		return nil, ErrorServer
	}

	return actList, SuccessCode
}

func RetrieveActivityListNotEndedByHost(id uint, pre int) (interface{}, uint) {
	actList := &dao.ActivityDigestArr{}
	err := actList.RetrieveNotEndedByHost(id, pre, 10)
	if err != nil {
		model.ErrLog.Println(err)
		return nil, ErrorServer
	}

	return actList, SuccessCode
}

func RetrieveActivityListRecommend(pre int) (interface{}, uint) {
	actList := &dao.ActivityDigestArr{}
	err := actList.RetrieveNotStart(pre, 10)
	if err != nil {
		model.ErrLog.Println(err)
		return nil, ErrorServer
	}

	return actList, SuccessCode
}

/*
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

*/
