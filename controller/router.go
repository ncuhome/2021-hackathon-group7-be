package controller

import (
	"github.com/gin-gonic/gin"
	"time"
)

func Run() {
	// 注意，handler使用同一个limit，会共同受到限制
	minute40 := limitIP(time.Minute, 40)
	minute2 := limitIP(time.Minute, 2)
	hour30 := limitIP(time.Hour, 30)

	g := gin.New()
	g.Use(gin.Logger(), gin.Recovery(), cors)

	g.GET("/user-info", minute40, GetUserInfo)
	g.GET("/activity/comment", minute40, GetCommentByActivity)
	g.GET("/user/comment", minute40, GetCommentByUser)
	g.GET("/verification/user", minute40, GetUserByV)

	g.POST("/register", hour30, Register)
	g.POST("/login", hour30, Login)
	g.POST("/email/password-key", minute2, SendPasswordEmailKey)
	g.POST("/email/password", minute40, SetPasswordByEmail)
	g.POST("/picture", minute40, PostPicture)

	g.GET("/activities/all", minute40, GetAllActivities)
	g.GET("/activities/detail", minute40, GetActivity)
	g.GET("/activities/place", minute40, GetActivitiesByPlace)
	g.GET("/activities/host", minute40, GetActivitiesByHost)

	a := g.Group("/auth", token)

	a.GET("/token", minute40, Verify)

	a.POST("/password", minute40, SetPassword)
	a.POST("/email/binding-key", minute2, SendBindEmailKey)
	a.POST("/email/binding", minute40, BindEmail)
	a.POST("/comment", minute40, PostComment)

	a.PUT("/user-info", minute40, PutUserInfo)
	a.PUT("/verification", minute40, PutV)

	a.DELETE("/email/binding", minute40, RemoveEmail)

	a.POST("/activity", minute40, CreateActivity)

	err := g.Run(":" + GinConfigObj.Port)
	if err != nil {
		panic(err)
	}
	return
}
