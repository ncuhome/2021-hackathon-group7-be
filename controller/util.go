package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

var HttpStatus = map[uint]int{
	1: 500,
	2: 400,
	3: 401,
}

var Message = map[uint]string{
	1:  "服务端错误",
	2:  "提交的数据错误",
	3:  "登录状态无效，请重新登录",
	4:  "用户名不符合要求",
	5:  "密码不符合要求",
	6:  "用户名已被注册",
	7:  "用户或密码错误",
	8:  "已经绑定了邮箱",
	9:  "邮件发送失败",
	10: "验证码不存在或过期",
	11: "验证码错误",
	12: "验证码与邮箱不匹配",
	13: "邮箱已被其它用户绑定",
	14: "旧密码错误",
	15: "该邮箱未绑定任何用户",
	16: "设置密码凭证不存在或已过期",
	17: "请求频繁，稍后重试",
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
		c.JSON(HttpStatus[0], map[string]interface{}{
			"code":    0,
			"message": "成功",
		})

	} else {
		c.JSON(HttpStatus[0], map[string]interface{}{
			"code":    0,
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
