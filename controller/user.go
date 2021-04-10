package controller

import (
	"github.com/gin-gonic/gin"
	"nspyf/model/dto"
	"nspyf/service"
)

func Register(c *gin.Context) {
	req := &dto.Register{}
	err := c.ShouldBind(req)
	if err != nil {
		RespondError(c, 2)
		return
	}

	code := service.Register(req)
	if code != 0 {
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
		RespondError(c, 2)
		return
	}

	data, code := service.Login(req)
	if code != 0 {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, data)
	return
}

func Verify(c *gin.Context) {
	_, err := GetClaimsSubAsID(c)
	if err != nil {
		RespondError(c, 3)
		return
	}

	RespondSuccess(c, nil)
	return
}

func SendBindEmailKey(c *gin.Context) {
	req := &dto.Email{}
	err := c.ShouldBind(req)
	if err != nil {
		RespondError(c, 2)
		return
	}

	id, err := GetClaimsSubAsID(c)
	if err != nil {
		RespondError(c, 3)
		return
	}

	code := service.SendBindEmailKey(req, id)
	if code != 0 {
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
		RespondError(c, 2)
		return
	}

	id, err := GetClaimsSubAsID(c)
	if err != nil {
		RespondError(c, 3)
		return
	}

	code := service.BindEmail(req, id)
	if code != 0 {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, nil)
	return
}

func RemoveEmail(c *gin.Context) {
	id, err := GetClaimsSubAsID(c)
	if err != nil {
		RespondError(c, 3)
		return
	}

	code := service.RemoveEmail(id)
	if code != 0 {
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
		RespondError(c, 2)
		return
	}

	id, err := GetClaimsSubAsID(c)
	if err != nil {
		RespondError(c, 3)
		return
	}

	code := service.SetPassword(req, id)
	if code != 0 {
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
		RespondError(c, 2)
		return
	}

	code := service.SendPasswordEmailKey(req)
	if code != 0 {
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
		RespondError(c, 2)
		return
	}

	code := service.SetPasswordByEmail(req)
	if code != 0 {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, nil)
	return
}
