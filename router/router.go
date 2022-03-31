package router

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"time"
	"tudo/controller"
	_ "tudo/docs"
	"tudo/putable"
)

func Run() {
	// 注意，handler使用同一个limit，会共同受到限制
	//minute80 := controller.LimitIP(time.Minute, 80)
	//minute4 := controller.LimitIP(time.Minute, 4)
	//hour30 := controller.LimitIP(time.Hour, 30)

	// test
	minute80 := controller.LimitIP(time.Minute, 80000)
	hour30 := controller.LimitIP(time.Hour, 30000)

	g := gin.New()
	g.Use(gin.Logger(), gin.Recovery(), controller.Cors)

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	g.GET("/user-info", minute80, controller.GetUserInfo)
	//g.GET("/verification/user", minute80, controller.GetUserByV)

	//g.GET("/activities/all", minute80, controller.GetAllActivities)
	//g.GET("/activities/detail", minute80, controller.GetActivity)
	//g.GET("/activities/place", minute80, controller.GetActivitiesByPlace)
	//g.GET("/activities/host", minute80, controller.GetActivitiesByHost)
	g.GET("/activity", minute80, controller.RetrieveActivity)
	g.GET("/not-start-activity", minute80, controller.RetrieveActivityNotStart)
	g.GET("/during-activity", minute80, controller.RetrieveActivityDuring)
	g.GET("/ended-activity", minute80, controller.RetrieveActivityEnded)
	g.GET("/recommend-activity", minute80, controller.RetrieveActivityRecommend)

	g.POST("/login", hour30, controller.Login)
	g.POST("/login/ncuos-token", hour30, controller.NCUOSTokenLogin)

	a := g.Group("/auth", controller.Token)

	a.POST("/picture", minute80, controller.PostPicture)

	a.GET("/token", minute80, controller.Verify)
	//a.GET("/email", minute80, controller.GetEmail)
	a.GET("/organization", minute80, controller.GetLeaderOrg)
	a.GET("/org/not-ended-activity", minute80, controller.RetrieveActivityNotEndedByHost)
	a.GET("/org/ended-activity", minute80, controller.RetrieveActivityEndedByHost)

	//a.POST("/password", minute80, controller.SetPassword)
	a.POST("/activity", minute80, controller.CreateActivity)
	a.POST("/organization", hour30, controller.OrgPostInfo)

	//a.PUT("/user-info", minute80, controller.PutUserInfo)
	a.PUT("/activity", minute80, controller.UpdateActivity)

	a.DELETE("/activity", minute80, controller.DeleteActivity)

	//desc 运维api
	g.POST("admin/login", putable.Login)
	b := g.Group("/put-table", putable.Jwt)
	b.GET("get-table", putable.GetTable)
	b.POST("update-table", putable.UpdateTable)
	b.POST("add-table", putable.AddTable)

	err := g.Run(":21001")
	if err != nil {
		panic(err)
	}
	return
}
