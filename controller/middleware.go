package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"nspyf/model"
	"nspyf/model/dao"
	"strconv"
	"time"
)

func cors(c *gin.Context) { // 预检
	origin := c.Request.Header.Get("Origin")
	c.Header("Access-Control-Allow-Origin", origin)      // 允许cookie不能使用*，必须有明确的origin
	c.Header("Access-Control-Allow-Credentials", "true") // 允许cookie
	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With, Host, Token")
	c.Header("Access-Control-Allow-Methods", "OPTIONS, GET, POST, DELETE, PUT, PATCH")
	c.Header("Access-Control-Max-Age", "3600") // 预检请求的有效期/秒
	c.Header("Content-Type", "application/json; charset=utf-8")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent) // 204
		return
	}
	c.Next()
}

func token(c *gin.Context) { // token验证
	token := c.Request.Header.Get("Token")
	claims, err := model.Jwt.ParseToken(token)
	if err != nil {
		RespondError(c, 3)
		return
	}

	if claims == nil {
		RespondError(c, 3)
		return
	}

	//claims sub为用户id，claims id为loginStatus
	//loginStatus校验
	user := &dao.UserDao{}
	id, err := strconv.Atoi(claims.Subject)
	if err != nil {
		RespondError(c, 3)
		return
	}
	err = user.GetProfile(uint(id))
	if err != nil {
		RespondError(c, 3)
		return
	}
	if user.Profile.LoginStatus != claims.Id {
		RespondError(c, 3)
		return
	}

	c.Set("claimsSub", claims.Subject)
	c.Next()
	return
}

func limitIP(interval time.Duration, limit int) gin.HandlerFunc {
	limiter := new(model.Limiter)
	limiter.Init(interval)
	return func(c *gin.Context) {
		ip := c.ClientIP()
		if limiter.LogAndCheck(ip, limit) == false {
			message := fmt.Sprintf("请求频繁，请%v后重试", interval.String())
			RespondErrorWith(c, 17, message)
			return
		}
		c.Next()
	}
}
