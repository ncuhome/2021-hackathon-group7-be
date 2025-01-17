package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"tudo/service"
)

var HttpStatus = map[uint]int{
	service.ErrorServer:     500,
	service.ErrorCommitData: 400,
	service.ErrorToken:      401,
}

var Message = map[uint]string{
	service.ErrorServer:           "服务端错误",
	service.ErrorCommitData:       "提交的数据错误",
	service.ErrorToken:            "无权访问",
	service.ErrorUsernameRepeated: "用户名已经被注册",
	service.ErrorLogin:            "用户名或密码错误",
	service.ErrorRequestRate:      "请求频繁，稍后重试",
}

func GetClaimsSubAsID(c *gin.Context) (uint, error) {
	sub, ok := c.Get("claimsSub")
	if ok == false {
		return 0, errors.New("get claimsSub from context failed")
	}

	IDStr, ok := sub.(string)
	if ok == false {
		return 0, errors.New("claimsSub transform failed")
	}

	idInt, err := strconv.Atoi(IDStr)
	if err != nil {
		return 0, errors.New("claimsSub is not int")
	}

	return uint(idInt), nil
}

func RespondSuccess(c *gin.Context, data interface{}) {
	if data == nil {
		c.JSON(HttpStatus[service.SuccessCode], map[string]interface{}{
			"code":    service.SuccessCode,
			"message": "成功",
		})

	} else {
		c.JSON(HttpStatus[service.SuccessCode], map[string]interface{}{
			"code":    service.SuccessCode,
			"message": "成功",
			"data":    data,
		})
	}
	c.Abort()
	return
}

func RespondError(c *gin.Context, code uint) {
	httpStatus, ok := HttpStatus[code]
	if !ok {
		httpStatus = 403
	}
	message, ok := Message[code]
	if !ok {
		message = "未描述的错误"
	}
	c.JSON(httpStatus, map[string]interface{}{
		"code":    code,
		"message": message,
	})
	c.Abort()
}

func RespondErrorWith(c *gin.Context, code uint, message string) {
	httpStatus, ok := HttpStatus[code]
	if !ok {
		httpStatus = 403
	}
	c.JSON(httpStatus, map[string]interface{}{
		"code":    code,
		"message": message,
	})
	c.Abort()
}
