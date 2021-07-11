package controller

import (
	"github.com/gin-gonic/gin"
	"tudo/model/dto"
	"tudo/service"
)

func Register(c *gin.Context) {
	req := &dto.Register{}
	err := c.ShouldBind(req)
	if err != nil {
		RespondError(c, service.CommitDataError)
		return
	}

	code := service.Register(req)
	if code != service.SuccessCode {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, nil)
	return
}

func Login(c *gin.Context) {
	req := &dto.Login{}
	err := c.ShouldBind(req)
	if err != nil {
		RespondError(c, service.CommitDataError)
		return
	}

	data, code := service.Login(req)
	if code != service.SuccessCode {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, data)
	return
}

func NCUOSTokenLogin(c *gin.Context) {
	req := &dto.Token{}
	err := c.ShouldBind(req)
	if err != nil {
		RespondError(c, service.CommitDataError)
		return
	}

	data, code := service.NCUOSTokenLogin(req)
	if code != service.SuccessCode {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, data)
	return
}

func Verify(c *gin.Context) {
	_, err := GetClaimsSubAsID(c)
	if err != nil {
		RespondError(c, service.TokenError)
		return
	}

	RespondSuccess(c, nil)
	return
}

func SendRegisterEmailKey(c *gin.Context) {
	req := &dto.Email{}
	err := c.ShouldBind(req)
	if err != nil {
		RespondError(c, service.CommitDataError)
		return
	}

	code := service.SendRegisterEmailKey(req)
	if code != service.SuccessCode {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, nil)
	return
}

func SendBindEmailKey(c *gin.Context) {
	req := &dto.Email{}
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

	code := service.SendBindEmailKey(req, id)
	if code != service.SuccessCode {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, nil)
	return
}

func BindEmail(c *gin.Context) {
	req := &dto.EmailBind{}
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

	code := service.BindEmail(req, id)
	if code != service.SuccessCode {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, nil)
	return
}

func RemoveEmail(c *gin.Context) {
	id, err := GetClaimsSubAsID(c)
	if err != nil {
		RespondError(c, service.TokenError)
		return
	}

	code := service.RemoveEmail(id)
	if code != service.SuccessCode {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, nil)
	return
}

func SetPassword(c *gin.Context) {
	req := &dto.SetPassword{}
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

	code := service.SetPassword(req, id)
	if code != service.SuccessCode {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, nil)
	return
}

func SendPasswordEmailKey(c *gin.Context) {
	req := &dto.Email{}
	err := c.ShouldBind(req)
	if err != nil {
		RespondError(c, service.CommitDataError)
		return
	}

	code := service.SendPasswordEmailKey(req)
	if code != service.SuccessCode {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, nil)
	return
}

func SetPasswordByEmail(c *gin.Context) {
	req := &dto.SetPasswordByEmailReq{}
	err := c.ShouldBind(req)
	if err != nil {
		RespondError(c, service.CommitDataError)
		return
	}

	code := service.SetPasswordByEmail(req)
	if code != service.SuccessCode {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, nil)
	return
}

func GetEmail(c *gin.Context) {
	id, err := GetClaimsSubAsID(c)
	if err != nil {
		RespondError(c, service.TokenError)
		return
	}

	data, code := service.GetEmail(id)
	if code != service.SuccessCode {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, data)
	return
}