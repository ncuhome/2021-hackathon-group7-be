package router

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"time"
	"tudo/controller"
	_ "tudo/docs"
	"tudo/model"
	"tudo/model/dao"
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
	g.GET("/activity", minute80, controller.RetrieveActivity)                   //descp 接收 活动id ,返回 活动
	g.GET("/not-start-activity", minute80, controller.RetrieveActivityNotStart) //descp 接收 时间 ,返回活动列表
	g.GET("/during-activity", minute80, controller.RetrieveActivityDuring)
	g.GET("/ended-activity", minute80, controller.RetrieveActivityEnded)
	g.GET("/recommend-activity", minute80, controller.RetrieveActivityRecommend)

	g.POST("/login", hour30, controller.Login)                       //descp 接收 用户学号/社团账号 密码 .返回 登录信息 token
	g.POST("/login/ncuos-token", hour30, controller.NCUOSTokenLogin) //descp 接收ncu-os的token ,返回 新token

	a := g.Group("/auth", controller.Token)

	a.POST("/picture", minute80, controller.PostPicture) // descp 上传图片并返回图片地址

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
	g.POST("admin/login", dao.Login)
	b := g.Group("/put-table", model.PutableJwt)
	b.GET("get-table", dao.GetMap)
	b.POST("update-table", dao.UpdateMap)

	err := g.Run(":21001")
	if err != nil {
		panic(err)
	}
	return
}
