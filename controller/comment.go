package controller

/*
import (
	"github.com/gin-gonic/gin"
	"strconv"
	"tudo/model/dto"
	"tudo/service"
)

func PostComment(c *gin.Context) {
	req := &dto.PostComment{}
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

	code := service.PostComment(req, id)
	if code != service.SuccessCode {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, nil)
	return
}

func GetCommentByActivity(c *gin.Context) {
	idInt, err := strconv.Atoi(c.Query("id"))
	if err != nil || idInt <= 0 {
		RespondError(c, service.ErrorCommitData)
		return
	}

	preInt, err := strconv.Atoi(c.DefaultQuery("pre", "0"))
	if err != nil || preInt < 0 {
		RespondError(c, service.ErrorCommitData)
		return
	}

	data, code := service.GetCommentByActivity(uint(idInt), uint(preInt))
	if code != service.SuccessCode {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, data)
	return
}

func GetCommentByUser(c *gin.Context) {
	idInt, err := strconv.Atoi(c.Query("id"))
	if err != nil || idInt <= 0 {
		RespondError(c, service.ErrorCommitData)
		return
	}

	preInt, err := strconv.Atoi(c.DefaultQuery("pre", "0"))
	if err != nil || preInt < 0 {
		RespondError(c, service.ErrorCommitData)
		return
	}

	data, code := service.GetCommentByUser(uint(idInt), uint(preInt))
	if code != service.SuccessCode {
		RespondError(c, code)
		return
	}

	RespondSuccess(c, data)
	return
}


*/
