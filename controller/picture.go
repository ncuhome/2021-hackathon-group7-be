package controller

import (
	"github.com/gin-gonic/gin"
	"nspyf/service"
)

func PostPicture(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		RespondError(c, service.CommitDataError)
	}

	data, code := service.PostPicture(file, header.Filename)
	if code != service.SuccessCode {
		RespondError(c, code)
	}

	RespondSuccess(c, data)
	return
}
