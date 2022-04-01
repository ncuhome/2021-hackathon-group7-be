package putable

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"tudo/model/dao"
)

type User struct { // descp 运维人员 tip 手动录入信息到数据库
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Leader struct { // descp 社团负责人
	Code         int    // descp 用于标记 可更改信息
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
		Width:       200,
		Editable:    true,
		Align:       "center",
		HeaderAlign: "center",
	}, {
		HeaderName:  "leader-name",
		Type:        "string",
		Width:       200,
		Editable:    true,
		Align:       "center",
		HeaderAlign: "center",
	}}
)

func Login(c *gin.Context) {
	var user User
	err := c.ShouldBind(&user)
	if err != nil {
		log.Println(err)
		return
	}

	if dao.DB.Model(&User{}).First(&user).Error != nil {
		ErrResponse(c, "用户名或密码错误")
		return
	}

	token := GenerateToken(user.Account)
	fmt.Println(token)
	if token == "" {
		ErrResponse(c, "token生成失败")
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

	for _, temp := range tables { // descp 转一下 []string
		data.Table = append(data.Table, []string{temp.Organization, temp.LeaderName})
	}

	c.AsciiJSON(200, gin.H{
		"data": data,
	})

}

func UpdateTable(c *gin.Context) {

	tx := dao.DB.Begin()
	defer func() {
		if recover() != nil {
			tx.Rollback()
			panic(recover())
		}
	}()

	if tx.Error != nil {
		ErrResponse(c, "errServer")
		return
	}

	var tables struct {
		Table [][]string `json:"table"`
	}
	err := c.BindJSON(&tables)
	if err != nil {
		ErrResponse(c, "ErrServer")

		return
	}

	leaders := make([]Leader, len(tables.Table))
	for i, temp := range tables.Table {
		leaders[i].LeaderName = temp[0]
		leaders[i].Organization = temp[1]
		leaders[i].Code = 0
	}

	// tip 直接全删除再全写入
	err = dao.DB.Delete(&User{}, "code = ?", '0').Error
	if err != nil {
		tx.Rollback()
		ErrResponse(c, "更改失败")
		return
	}
	err = dao.DB.Model(&Leader{}).Create(&leaders).Error
	if err != nil {
		tx.Rollback()
		ErrResponse(c, "更改失败")
		return
	}

	c.AsciiJSON(200, gin.H{
		"message": "success",
	})

}

func ErrResponse(c *gin.Context, msg string) {
	c.AsciiJSON(400, gin.H{
		"message": msg,
	})
}
