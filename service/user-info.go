package service

import (
	"gorm.io/gorm"
	"nspyf/model"
	"nspyf/model/dao"
	"nspyf/model/dto"
)

func PutUserInfo(req *dto.UserInfo, id uint) (*map[string]interface{}, uint) {
	userInfo := &dao.UserInfo{
		UserID: id,
	}
	newInfo := &map[string]interface{}{
		"nickname":req.Nickname,
		"avatar": req.Avatar,
		"digest":req.Digest,
	}
	err := userInfo.Update(newInfo)
	if err != nil {
		model.ErrLog.Println(err)
		return nil, ServerError
	}
	return newInfo, SuccessCode
}

func GetUserInfo(id uint) (*map[string]interface{}, uint) {
	userInfo := &dao.UserInfo{
		UserID: id,
	}

	err := userInfo.Retrieve()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, CommitDataError
		}
		model.ErrLog.Println(err)
		return nil, ServerError
	}

	data := &map[string]interface{}{
		"nickname":userInfo.Nickname,
		"avatar": userInfo.Avatar,
		"digest":userInfo.Digest,
		"verification":userInfo.Verification,
	}
	return data, SuccessCode
}