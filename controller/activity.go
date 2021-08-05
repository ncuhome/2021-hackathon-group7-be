package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"tudo/model/dto"
	"tudo/service"
)

// @Summary 社团账号发布活动
// @Tags 活动系统
// @Accept application/json
// @Produce application/json
// @Param Token header string true "用户令牌"
// @Param JSON body dto.Activity true " "
// @Router /auth/activity [post]
func CreateActivity(c *gin.Context) {
	req := &dto.Activity{}
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

	code := service.CreateActivity(req, id)
	if code != 0 {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, nil)
	return
}

// @Summary 社团账号修改活动
// @Tags 活动系统
// @Accept application/json
// @Produce application/json
// @Param Token header string true "用户令牌"
// @Param Query query string true "id"
// @Param JSON body dto.Activity true " "
// @Router /auth/activity [put]
func UpdateActivity(c *gin.Context) {
	actIDInt, err := strconv.Atoi(c.Query("id"))
	if err != nil || actIDInt <= 0 {
		RespondError(c, service.ErrorCommitData)
		return
	}

	req := &dto.Activity{}
	err = c.ShouldBind(req)
	if err != nil {
		RespondError(c, service.ErrorCommitData)
		return
	}

	id, err := GetClaimsSubAsID(c)
	if err != nil {
		RespondError(c, service.ErrorToken)
		return
	}

	code := service.UpdateActivity(req, uint(actIDInt), id)
	if code != 0 {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, nil)
	return
}

// @Summary 社团账号删除活动
// @Tags 活动系统
// @Accept application/json
// @Produce application/json
// @Param Token header string true "用户令牌"
// @Param Query query string true "id"
// @Router /auth/activity [delete]
func DeleteActivity(c *gin.Context) {
	idInt, err := strconv.Atoi(c.Query("id"))
	if err != nil || idInt <= 0 {
		RespondError(c, service.ErrorCommitData)
		return
	}

	id, err := GetClaimsSubAsID(c)
	if err != nil {
		RespondError(c, service.ErrorToken)
		return
	}

	code := service.DeleteActivity(uint(idInt), id)
	if code != 0 {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, nil)
	return
}

/*
func GetAllActivities(c *gin.Context) {
	respond, code := service.GetAllActivities()

	if code != 0 {
		RespondError(c, code)
		return
	}
	RespondSuccess(c, respond)
	return
}

func GetActivity(c *gin.Context) {
	id := c.PostForm("id")
	respond, code := service.GetActivity(id)

	if code != 0 {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, respond)
	return
}

func GetActivitiesByPlace(c *gin.Context) {
	place := c.PostForm("place")
	respond, code := service.GetActivitiesByPlace(place)
	if code != 0 {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, respond)
	return
}

func GetActivitiesByHost(c *gin.Context) {
	hostInt, err := strconv.Atoi(c.PostForm("host"))
	if err != nil || hostInt <= 0 {
		RespondError(c, service.ErrorCommitData)
		return
	}

	respond, code := service.GetActivitiesByHost(uint(hostInt))
	if code != 0 {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, respond)
	return
}


 */