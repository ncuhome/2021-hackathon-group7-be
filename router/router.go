package router

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"time"
	"tudo/controller"
	_ "tudo/docs"
)

func Run() {
	// 注意，handler使用同一个limit，会共同受到限制
	minute80 := controller.LimitIP(time.Minute, 80)
	//minute4 := controller.LimitIP(time.Minute, 4)
	hour30 := controller.LimitIP(time.Hour, 30)

	g := gin.New()
	g.Use(gin.Logger(), gin.Recovery(), controller.Cors)

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//g.GET("/user-info", minute80, controller.GetUserInfo)
	//g.GET("/activity/comment", minute80, controller.GetCommentByActivity)
	//g.GET("/user/comment", minute80, controller.GetCommentByUser)
	//g.GET("/verification/user", minute80, controller.GetUserByV)

	//g.GET("/activities/all", minute80, controller.GetAllActivities)
	//g.GET("/activities/detail", minute80, controller.GetActivity)
	//g.GET("/activities/place", minute80, controller.GetActivitiesByPlace)
	//g.GET("/activities/host", minute80, controller.GetActivitiesByHost)

	//g.POST("/register/email-key", minute4, controller.SendRegisterEmailKey)
	g.POST("/login", hour30, controller.Login)
	g.POST("/login/ncuos-token", hour30, controller.NCUOSTokenLogin)
	//g.POST("/email/password-key", minute4, controller.SendPasswordEmailKey)
	//g.POST("/email/password", minute80, controller.SetPasswordByEmail)

	a := g.Group("/auth", controller.Token)

	//a.POST("/picture", minute80, controller.PostPicture)

	a.GET("/token", minute80, controller.Verify)
	//a.GET("/email", minute80, controller.GetEmail)
	a.GET("/organization", minute80, controller.GetLeaderOrg)

	//a.POST("/password", minute80, controller.SetPassword)
	//a.POST("/email/binding-key", minute4, controller.SendBindEmailKey)
	//a.POST("/email/binding", minute80, controller.BindEmail)
	//a.POST("/comment", minute80, controller.PostComment)
	//a.POST("/activity", minute80, controller.CreateActivity)
	a.POST("/organization", hour30, controller.OrgPostInfo)

	//a.PUT("/user-info", minute80, controller.PutUserInfo)
	//a.PUT("/verification", minute80, controller.PutV)

	//a.DELETE("/email/binding", minute80, controller.RemoveEmail)

	err := g.Run(":21001")
	if err != nil {
		panic(err)
	}
	return
}
