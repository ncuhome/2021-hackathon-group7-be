package main

import (
	"github.com/gin-gonic/gin"
	"tudo/model"
	"tudo/model/dao"
	"tudo/router"
	"tudo/service"
)

// @title Swagger API
// @version 1.0
// @description 给出了请求方法，点击Model可以查看请求体模型及备注。
// @description 查看响应体需要打开浏览器开发者工具，在页面接口初Try it out，填写数据然后Execute。
// @contact.email 316851756@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host tudo-api.nspyf.top
// host localhost:21001

func main() {
	dao.DBInit("./config/db.json")
	dao.CacheInit("./config/cache.json")
	model.JwtInit("./config/jwt.json")
	model.OssInit()
	model.LogInit()
	service.SyncTencentDoc()

	gin.SetMode(gin.ReleaseMode)
	router.Run()
	return
}
