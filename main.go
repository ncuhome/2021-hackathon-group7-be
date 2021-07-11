package main

import (
	"github.com/gin-gonic/gin"
	"tudo/controller"
	"tudo/model"
	"tudo/model/dao"
	"tudo/root"
)

func main() {
	dao.DBInit("./config/db.json")
	dao.CacheInit("./config/cache.json")
	controller.GinInit("./config/gin.json")
	model.JwtInit("./config/jwt.json")
	model.EmailInit("./config/email.json")
	model.AdminInit("./config/admin.json")
	//model.PictureInit("/etc/share/nginx/html/picture/")
	model.OssInit()
	model.LogInit()

	gin.SetMode(gin.ReleaseMode)
	root.Run()
	return
}
