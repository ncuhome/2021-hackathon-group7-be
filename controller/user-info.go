package controller

import (
	"github.com/gin-gonic/gin"
	"nspyf/model/dto"
	"nspyf/service"
)

func PutUserInfo(c *gin.Context) {
	req := &dto.UserInfo{}
	err := c.ShouldBind(req)
	if err != nil {
		RespondError(c, service.CommitDataError)
		return
	}

	id, err := GetClaimsSubAsID(c)
	if err != nil {
		RespondError(c, service.TokenError)
		return
	}

	data, code := service.PutUserInfo(req, id)
	if code != service.SuccessCode {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, data)
	return
}

