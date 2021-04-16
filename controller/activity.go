package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nspyf/model/dto"
	"nspyf/service"
)

func CreateActivity(c *gin.Context) {
	req := &dto.Activities{}
	err := c.ShouldBind(req)

	if err != nil {
		RespondError(c, 11)
		return
	}

	code := service.CreateActivity(req)

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
	fmt.Println(id, 5)
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
	host := c.PostForm("host")
	respond, code := service.GetActivitiesByHost(host)
	if code != 0 {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, respond)
	return
}
