package service

import (
	"gorm.io/gorm"
	"tudo/model"
	"tudo/model/dao"
	"tudo/model/dto"
)

func PutUserInfo(req *dto.UserInfo, id uint) (*map[string]interface{}, uint) {
	userInfo := &dao.UserInfo{
		UserID: id,
	}
	newInfo := &map[string]interface{}{
		"nickname": req.Nickname,
		"avatar":   req.Avatar,
		"digest":   req.Digest,
	}
	err := userInfo.Update(newInfo)
	if err != nil {
		model.ErrLog.Println(err)
		return nil, ErrorServer
	}
	return newInfo, SuccessCode
}

func GetUserInfo(id uint) (*map[string]interface{}, uint) {
	userInfoDao := &dao.UserInfoDao{}

	err := userInfoDao.GetProfile(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrorCommitData
		}
		model.ErrLog.Println(err)
		return nil, ErrorServer
	}

	data := &map[string]interface{}{
		"nickname":     userInfoDao.Profile.Nickname,
		"avatar":       userInfoDao.Profile.Avatar,
		"digest":       userInfoDao.Profile.Digest,
		"verification": userInfoDao.Profile.Verification,
	}
	return data, SuccessCode
}

func GetLeaderOrg(id uint) (*map[string]interface{}, uint) {
	ncuUser := &dao.User{ID: id}
	err := ncuUser.Retrieve()
	if err != nil {
		return nil, ErrorCommitData
	}

	org := dao.LeaderMap[ncuUser.Phone].Organization
	data := &map[string]interface{}{
		"organization": org,
	}
	return data, SuccessCode
}

func GetUserByV(pre uint) (*map[string]interface{}, uint) {
	list := &dao.UserInfoProfileList{}
	err := list.RetrieveByV(pre)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, ErrorServer
		}
	}

	data := &map[string]interface{}{
		"list": list.Data,
	}
	return data, SuccessCode
}
