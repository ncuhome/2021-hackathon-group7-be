package service

import (
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