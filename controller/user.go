package controller

import (
	"github.com/gin-gonic/gin"
	"tudo/model/dto"
	"tudo/service"
)

// @Summary 云家园账号或社团账号登录
// @Tags 用户系统
// @Accept application/json
// @Produce application/json
// @Param JSON body dto.Login true " "
// @Router /login [post]
func Login(c *gin.Context) {
	req := &dto.Login{}
	err := c.ShouldBind(req)
	if err != nil {
		RespondError(c, service.ErrorCommitData)
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
// @Accept application/json
// @Produce application/json
// @Param JSON body dto.Token true " "
// @Router /login/ncuos-token [post]
func NCUOSTokenLogin(c *gin.Context) {
	req := &dto.Token{}
	err := c.ShouldBind(req)
	if err != nil {
		RespondError(c, service.ErrorCommitData)
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

// @Summary 检验token是否有效并返回用户角色
// @Tags 用户系统
// @Accept application/json
// @Produce application/json
// @Param Token header string true "用户令牌"
// @Router /auth/token [get]
func Verify(c *gin.Context) {
	id, err := GetClaimsSubAsID(c)
	if err != nil {
		RespondError(c, service.ErrorToken)
		return
	}

	data, code := service.GetRole(id)
	if(code != service.SuccessCode) {
		RespondError(c, code)
		return 
	}
	
	RespondSuccess(c, data)
	return
}

// @Summary 社团账号激活（注册）、修改
// @Tags 用户系统
// @Accept application/json
// @Produce application/json
// @Param Token header string true "用户令牌"
// @Param JSON body dto.OrgInfo true " "
// @Router /auth/organization [post]
func OrgPostInfo(c *gin.Context) {
	req := &dto.OrgInfo{}
	err := c.ShouldBind(req)
	if err != nil {
		RespondError(c, service.ErrorCommitData)
		return
	}

	id, err := GetClaimsSubAsID(c)
	if err != nil {
		RespondError(c, service.ErrorToken)
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