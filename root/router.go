package root

import (
	"github.com/gin-gonic/gin"
	"time"
	"tudo/controller"
)

func Run() {
	// 注意，handler使用同一个limit，会共同受到限制
	minute80 := controller.LimitIP(time.Minute, 80)
	minute4 := controller.LimitIP(time.Minute, 4)
	hour30 := controller.LimitIP(time.Hour, 30)

	g := gin.New()
	g.Use(gin.Logger(), gin.Recovery(), controller.Cors)

	g.GET("/user-info", minute80, controller.GetUserInfo)
	g.GET("/activity/comment", minute80, controller.GetCommentByActivity)
	g.GET("/user/comment", minute80, controller.GetCommentByUser)
	g.GET("/verification/user", minute80, controller.GetUserByV)

	g.POST("/register", hour30, controller.Register)
	g.POST("/register/email-key", minute4, controller.SendRegisterEmailKey)
	g.POST("/login", hour30, controller.Login)
	g.POST("/email/password-key", minute4, controller.SendPasswordEmailKey)
	g.POST("/email/password", minute80, controller.SetPasswordByEmail)
	g.POST("/picture", minute80, controller.PostPicture)

	g.GET("/activities/all", minute80, controller.GetAllActivities)
	g.GET("/activities/detail", minute80, controller.GetActivity)
	g.GET("/activities/place", minute80, controller.GetActivitiesByPlace)
	g.GET("/activities/host", minute80, controller.GetActivitiesByHost)

	a := g.Group("/auth", controller.Token)

	a.GET("/token", minute80, controller.Verify)
	a.GET("/email", minute80, controller.GetEmail)

	a.POST("/password", minute80, controller.SetPassword)
	a.POST("/email/binding-key", minute4, controller.SendBindEmailKey)
	a.POST("/email/binding", minute80, controller.BindEmail)
	a.POST("/comment", minute80, controller.PostComment)

	a.PUT("/user-info", minute80, controller.PutUserInfo)
	a.PUT("/verification", minute80, controller.PutV)

	a.DELETE("/email/binding", minute80, controller.RemoveEmail)

	a.POST("/activity", minute80, controller.CreateActivity)

	err := g.Run(":" + controller.GinConfigObj.Port)
	if err != nil {
		panic(err)
	}
	return
}
