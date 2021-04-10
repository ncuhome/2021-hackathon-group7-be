package main

import (
	"nspyf/controller"
	"nspyf/model"
	"nspyf/model/dao"
)

func main() {
	dao.DBInit("./config/db.json")
	dao.CacheInit("./config/cache.json")
	model.JwtInit("./config/jwt.json")
	model.EmailInit("./config/email.json")
	controller.GinInit("./config/gin.json")

	controller.Run()
	return
}
