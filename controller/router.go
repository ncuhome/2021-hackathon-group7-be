package controller

import (
	"github.com/gin-gonic/gin"
	"time"
)

func Run() {
	// 注意，handler使用同一个limit，会共同受到限制
	minute20 := limitIP(time.Minute, 20)
	minute2 := limitIP(time.Minute, 2)
	hour30 := limitIP(time.Hour, 30)

	g := gin.New()
	g.Use(gin.Logger(), gin.Recovery(), cors)

	g.GET("/user-info", minute20, GetUserInfo)
	g.GET("/activity/comment", minute20, GetCommentByActivity)
	g.GET("/user/comment", minute20, GetCommentByUser)

	g.POST("/register", hour30, Register)
	g.POST("/login", hour30, Login)
	g.POST("/email/password-key", minute2, SendPasswordEmailKey)
	g.POST("/email/password", minute20, SetPasswordByEmail)
	g.POST("/picture", minute20, PostPicture)

	g.GET("/activities/all", minute20, GetAllActivities)
	g.GET("/activities/detail", minute20, GetActivity)
	g.GET("/activities/place", minute20, GetActivitiesByPlace)
	g.GET("/activities/host", minute20, GetActivitiesByHost)

	a := g.Group("/auth", token)

	a.GET("/token", minute20, Verify)

	a.POST("/password", minute20, SetPassword)
	a.POST("/email/binding-key", minute2, SendBindEmailKey)
	a.POST("/email/binding", minute20, BindEmail)
	a.POST("/comment", minute20, PostComment)

	a.PUT("/user-info", minute20, PutUserInfo)
	a.PUT("/verification", minute20, PutV)

	a.DELETE("/email/binding", minute20, RemoveEmail)

	a.POST("/activity", minute20, CreateActivity)

	err := g.Run(":" + GinConfigObj.Port)
	if err != nil {
		panic(err)
	}
	return
}
