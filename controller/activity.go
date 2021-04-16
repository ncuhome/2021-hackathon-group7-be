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
		RespondError(c, service.CommitDataError)
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
