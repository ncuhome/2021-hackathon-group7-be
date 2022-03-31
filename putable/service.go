package putable

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"tudo/model/dao"
)

type User struct { // desc 运维人员 tip 手动录入信息到数据库
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Leader struct { // desc 社团负责人
	Organization string `json:"organization"`
	LeaderName   string `json:"leader-name"`
}

type column struct {
	HeaderName  string `json:"headerName"`  //表头
	Type        string `json:"type"`        //数据类型，暂时只支持string
	Width       int    `json:"width"`       //在表格中列的显示宽度
	Editable    bool   `json:"editable"`    //是否可编辑
	Align       string `json:"align"`       // 列对齐方式，left, right, center
	HeaderAlign string `json:"headerAlign"` //标题对齐方式，left, right, center
}

type SentData struct {
	ColumnOption []column   `json:"column_options"`
	Table        [][]string `json:"table"`
}

var (
	columns = []column{{
		HeaderName:  "organization",
		Type:        "string",
		Width:       100,
		Editable:    false,
		Align:       "center",
		HeaderAlign: "center",
	}, {
		HeaderName:  "leader-name",
		Type:        "string",
		Width:       100,
		Editable:    false,
		Align:       "center",
		HeaderAlign: "center",
	}}
)

var LeaderMap = map[string]Leader{} // 电话：leader // question 这是拿来干嘛的?

func Login(c *gin.Context) {
	var user User
	err := c.ShouldBind(&user)
	if err != nil {
		log.Println(err)
		return
	}

	token := GenerateToken(user.Account)
	fmt.Println(token)
	if token == "" {
		c.AsciiJSON(400, gin.H{
			"message": "token生成失败",
		})
		return
	}

	m := make(map[string]string)
	m["token"] = token

	c.AsciiJSON(200, gin.H{
		"data": m,
	})

}

func GetTable(c *gin.Context) {
	var data SentData
	data.ColumnOption = columns
	var tables []Leader

	dao.DB.Model(&Leader{}).Find(&tables)

	for _, temp := range tables { // desc 转一下 []string
		data.Table = append(data.Table, []string{temp.Organization, temp.LeaderName})
	}

	c.AsciiJSON(200, gin.H{
		"data": data,
	})

}

func UpdateTable(c *gin.Context) {
	//var tables [][]string

}

func AddTable(c *gin.Context) {
	//todo
}
