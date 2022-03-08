package controller

import (
	"github.com/gin-gonic/gin"
	"tudo/service"
)

/*
func PutUserInfo(c *gin.Context) {
	req := &dto.UserInfo{}
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

	data, code := service.PutUserInfo(req, id)
	if code != service.SuccessCode {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, data)
	return
}
 */

/*

func GetUserInfo(c *gin.Context) {
	idInt, err := strconv.Atoi(c.Query("id"))
	if err != nil || idInt <= 0 {
		RespondError(c, service.ErrorCommitData)
		return
	}

	data, code := service.GetUserInfo(uint(idInt))
	if code != service.SuccessCode {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, data)
	return
}

 */

// @Summary 判断用户是否是某个组织的负责人
// @Tags 用户系统
// @Accept json
// @Produce application/json
// @Param Token header string true "用户令牌"
// @Router /auth/organization [get]
func GetLeaderOrg(c *gin.Context) {
	id, err := GetClaimsSubAsID(c)
	if err != nil {
		RespondError(c, service.ErrorToken)
		return
	}

	data, code := service.GetLeaderOrg(id)
	if code != service.SuccessCode {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, data)
	return
}

/*
func GetUserByV(c *gin.Context) {
	preInt, err := strconv.Atoi(c.DefaultQuery("pre", "0"))
	if err != nil || preInt < 0 {
		RespondError(c, service.ErrorCommitData)
		return
	}

	data, code := service.GetUserByV(uint(preInt))
	if code != service.SuccessCode {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, data)
	return
}

 */