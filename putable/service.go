package putable

import (
	"github.com/gin-gonic/gin"
	"log"
)

type User struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var user User
	err := c.ShouldBind(&user)
	if err != nil {
		log.Println(err)
		return
	}

	token := GenerateToken(user.Account)
	//fmt.Println(token)
	//if token == "" {
	//	fmt.Println("token生成失败")
	//	return
	//}

	m := make(map[string]string)
	m["token"] = token
	c.AsciiJSON(200, gin.H{
		"data": m,
	})

}

func GetData(c *gin.Context) {

}

func UpdateData(c *gin.Context) {

}
func AddData(c *gin.Context) {

}
