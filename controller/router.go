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

	g.POST("/register", hour30, Register)
	g.POST("/login", hour30, Login)
	g.POST("/email/password-key", minute2, SendPasswordEmailKey)
	g.POST("/email/password", minute20, SetPasswordByEmail)

	a := g.Group("/auth", token)

	a.GET("/token", minute20, Verify)

	a.POST("/password", minute20, SetPassword)
	a.POST("/email/binding-key", minute2, SendBindEmailKey)
	a.POST("/email/binding", minute20, BindEmail)

	a.PUT("/user-info", minute20, PutUserInfo)

	a.DELETE("/email/binding", minute20, RemoveEmail)

	err := g.Run(":" + GinConfigObj.Port)
	if err != nil {
		panic(err)
	}
	return
}
