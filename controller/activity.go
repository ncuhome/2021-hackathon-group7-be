package controller

import (
	"github.com/gin-gonic/gin"
	"nspyf/model/dto"
	"nspyf/service"
	"strconv"
)

func CreateActivity(c *gin.Context) {
	req := &dto.Activities{}
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

	code := service.CreateActivity(req, id)

	if code != 0 {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, nil)
	return
}

func GetAllActivities(c *gin.Context) {
	pre, err := strconv.Atoi(c.DefaultQuery("pre", "0"))
	if err != nil || pre < 0 {
		RespondError(c, service.CommitDataError)
		return
	}

	respond, code := service.GetAllActivities(pre)

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
	place := c.Query("place")
	pre, err := strconv.Atoi(c.DefaultQuery("pre", "0"))
	if err != nil || pre < 0 {
		RespondError(c, service.CommitDataError)
		return
	}
	respond, code := service.GetActivitiesByPlace(place, pre)
	if code != 0 {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, respond)
	return
}

func GetActivitiesByHost(c *gin.Context) {
	hostInt, err := strconv.Atoi(c.Query("host"))
	if err != nil || hostInt <= 0 {
		RespondError(c, service.CommitDataError)
		return
	}
	pre, err := strconv.Atoi(c.DefaultQuery("pre", "0"))
	if err != nil || pre < 0 {
		RespondError(c, service.CommitDataError)
		return
	}

	respond, code := service.GetActivitiesByHost(uint(hostInt), pre)
	if code != 0 {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, respond)
	return
}
