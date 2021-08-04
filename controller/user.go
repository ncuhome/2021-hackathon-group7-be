package controller

import (
	"github.com/gin-gonic/gin"
	"tudo/model/dto"
	"tudo/service"
)

// @Summary 云家园账号或社团账号登录
// @Tags 用户系统
// @Accept json
// @Produce application/json
// @Param object body dto.Login true " "
// @Router /login [post]
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

// @Summary 云家园账号token登录
// @Tags 用户系统
// @Accept json
// @Produce application/json
// @Param object body dto.Token true " "
// @Router /login/ncuos-token [post]
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

// @Summary 检验token是否有效
// @Tags 用户系统
// @Accept json
// @Produce application/json
// @Param Token header string true "用户令牌"
// @Router /auth/token [get]
func Verify(c *gin.Context) {
	_, err := GetClaimsSubAsID(c)
	if err != nil {
		RespondError(c, service.TokenError)
		return
	}

	RespondSuccess(c, nil)
	return
}

// @Summary 社团账号激活（注册）、修改
// @Tags 用户系统
// @Accept json
// @Produce application/json
// @Param Token header string true "用户令牌"
// @Param object body dto.OrgInfo true " "
// @Router /auth/organization [post]
func OrgPostInfo(c *gin.Context) {
	req := &dto.OrgInfo{}
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

	code := service.OrgPostInfo(req, id)
	if code != service.SuccessCode {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, nil)
	return
}

/*
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

*/