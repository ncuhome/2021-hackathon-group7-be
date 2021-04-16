package service

import (
	"gorm.io/gorm"
	"tudo/model/dao"
	"tudo/model/dto"
)

func PostComment(req *dto.PostComment, id uint) uint {
	comment := &dao.Comment{
		Content:    req.Content,
		UserID:     id,
		ActivityID: req.ActivityID,
	}

	// TODO 判断活动是否存在

	err := comment.Create()
	if err != nil {
		return ServerError
	}
	return SuccessCode
}

func GetCommentByActivity(id uint, pre uint) (*map[string]interface{}, uint) {
	commentList := &dao.CommentList{}
	err := commentList.RetrieveByActivity(id, pre)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, ServerError
		}
	}

	commentUsers := make([]*map[string]interface{},0)
	for _,v := range commentList.Data {
		userInfoDao := &dao.UserInfoDao{}
		err := userInfoDao.GetDigest(v.UserID)
		if err != nil {
			return nil, ServerError
		}

		commentUser := &map[string]interface{}{
			"comment": v,
			"user": userInfoDao.Digest,
		}
		commentUsers = append(commentUsers, commentUser)
	}

	data := &map[string]interface{}{
		"list": commentUsers,
	}
	return data, SuccessCode
}

func GetCommentByUser(id uint, pre uint) (*map[string]interface{}, uint) {
	commentList := &dao.CommentList{}
	err := commentList.RetrieveByUser(id, pre)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, ServerError
		}
	}

	commentUsers := make([]*map[string]interface{},0)
	for _,v := range commentList.Data {
		userInfoDao := &dao.UserInfoDao{}
		err := userInfoDao.GetDigest(v.UserID)
		if err != nil {
			return nil, ServerError
		}

		commentUser := &map[string]interface{}{
			"comment": v,
			"user": userInfoDao.Digest,
		}
		commentUsers = append(commentUsers, commentUser)
	}

	data := &map[string]interface{}{
		"list": commentUsers,
	}
	return data, SuccessCode
}
