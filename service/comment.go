package service

import (
	"gorm.io/gorm"
	"nspyf/model/dao"
	"nspyf/model/dto"
)

func PostComment(req *dto.PostComment, id uint) uint {
	comment := &dao.Comment{
		Content:        req.Content,
		UserID:         id,
		ActivityID:     req.ActivityID,
	}

	// TODO 判断活动是否存在

	err := comment.Create()
	if err != nil {
		return ServerError
	}
	return SuccessCode
}

func GetCommentByActivity(id uint,pre uint) (*map[string]interface{},uint) {
	commentList := &dao.CommentList{}
	err := commentList.RetrieveByActivity(id, pre)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, ServerError
		}
	}
	data := &map[string]interface{}{
		"data": commentList.Data,
	}
	return data, SuccessCode
}