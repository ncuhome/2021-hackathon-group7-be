package main

import (
	"github.com/gin-gonic/gin"
	"tudo/controller"
	"tudo/model"
	"tudo/model/dao"
	"tudo/router"
	"tudo/service"
)

func main() {
	dao.DBInit("./config/db.json")
	dao.CacheInit("./config/cache.json")
	controller.GinInit("./config/gin.json")
	model.JwtInit("./config/jwt.json")
	model.EmailInit("./config/email.json")
	model.AdminInit("./config/admin.json")
	model.OssInit()
	model.LogInit()
	service.SyncTencentDoc()

	gin.SetMode(gin.ReleaseMode)
	router.Run()
	return
}
