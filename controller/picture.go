package controller

import (
	"github.com/gin-gonic/gin"
	"tudo/service"
)

// @Summary 图片上传接口
// @Tags  其它
// @Accept multipart/form-data
// @Produce application/json
// @Param Token header string true "用户令牌"
// @Param formData body object true "file字段放图片数据"
// @Router /auth/picture [post]
func PostPicture(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		RespondError(c, service.ErrorCommitData)
		return
	}

	data, code := service.PostPicture(file, header.Filename)
	if code != service.SuccessCode {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, data)
	return
}
