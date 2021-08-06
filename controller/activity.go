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
// @Param id query string true "活动id"
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
// @Param id query string true "活动id"
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

// @Summary 获取活动详情
// @Tags 活动系统
// @Produce application/json
// @Param id query string true "活动id"
// @Router /activity [get]
func RetrieveActivity(c *gin.Context) {
	idInt, err := strconv.Atoi(c.Query("id"))
	if err != nil || idInt <= 0 {
		RespondError(c, service.ErrorCommitData)
		return
	}

	data, code := service.RetrieveActivity(uint(idInt))
	if code != 0 {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, data)
	return
}

// @Summary 分页获取未开始的活动列表
// @Tags 活动系统
// @Produce application/json
// @Param pre query string true "上一次调用本接口得到的活动列表的最后一个活动的开始时间戳，第一次调用用当前时间戳"
// @Router /not-start-activity [get]
func RetrieveActivityNotStart(c *gin.Context) {
	pre, err := strconv.Atoi(c.Query("pre"))
	if err != nil || pre <= 0 {
		RespondError(c, service.ErrorCommitData)
		return
	}

	data, code := service.RetrieveActivityListNotStart(pre)
	if code != 0 {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, data)
	return
}

// @Summary 分页获取正在进行的活动列表
// @Tags 活动系统
// @Produce application/json
// @Param pre query string true "上一次调用本接口得到的活动列表的最后一个活动的开始时间戳，第一次调用用当前时间戳"
// @Param now query string true "当前时间戳"
// @Router /during-activity [get]
func RetrieveActivityDuring(c *gin.Context) {
	now, err := strconv.Atoi(c.Query("now"))
	if err != nil || now <= 0 {
		RespondError(c, service.ErrorCommitData)
		return
	}

	pre, err := strconv.Atoi(c.Query("pre"))
	if err != nil || pre <= 0 {
		RespondError(c, service.ErrorCommitData)
		return
	}

	data, code := service.RetrieveActivityListDuring(now, pre)
	if code != 0 {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, data)
	return
}

// @Summary 分页获取历史活动列表
// @Tags 活动系统
// @Produce application/json
// @Param pre query string true "上一次调用本接口得到的活动列表的最后一个活动的结束时间戳，第一次调用用当前时间戳"
// @Router /ended-activity [get]
func RetrieveActivityEnded(c *gin.Context) {
	pre, err := strconv.Atoi(c.Query("pre"))
	if err != nil || pre <= 0 {
		RespondError(c, service.ErrorCommitData)
		return
	}

	data, code := service.RetrieveActivityListEnded(pre)
	if code != 0 {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, data)
	return
}

// @Summary 分页获取推荐活动列表
// @Tags 活动系统
// @Produce application/json
// @Param pre query string true "上一次调用本接口得到的活动列表的最后一个活动的开始时间戳，第一次调用用当前时间戳"
// @Router /recommend-activity [get]
func RetrieveActivityRecommend(c *gin.Context) {
	pre, err := strconv.Atoi(c.Query("pre"))
	if err != nil || pre <= 0 {
		RespondError(c, service.ErrorCommitData)
		return
	}

	//TODO 暂时用未进行活动，以后改其它
	data, code := service.RetrieveActivityListRecommend(pre)
	if code != 0 {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, data)
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