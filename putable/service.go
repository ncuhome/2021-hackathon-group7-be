package putable

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"tudo/model/dao"
	_ "tudo/model/dto"
)

// tip 使用须知
//  ①登录	/admin/longin
//  ②必须先get一下table		/put-table/get-table
//  ③再下载表格进行更改 或者直接传入完整的表格	/put-table/update-table

type User struct { // descp 运维人员 tip 手动录入信息到数据库
	Account  string `json:"account" binding:"required" gorm:"unique_index;not null"`
	Password string `json:"password" binding:"required"`
}

type Leader struct { // descp 社团负责人
	Phone        string `json:"phone"`
	Organization string `json:"organization"`
	LeaderName   string `json:"leader-name"`
}

var LeaderMap = map[string]Leader{} // 电话：leader

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
	}, {
		HeaderName:  "phone",
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

func GetMap(c *gin.Context) {
	var data SentData
	data.ColumnOption = columns
	for _, temp := range LeaderMap { // descp 转一下 []string
		data.Table = append(data.Table, []string{temp.Organization, temp.LeaderName, temp.Phone})
	}
	c.AsciiJSON(200, gin.H{
		"data": data,
	})
}

func UpdateMap(c *gin.Context) {

	var tables struct {
		Table [][]string `json:"table"`
	}
	err := c.ShouldBind(&tables)
	if err != nil {
		ErrResponse(c, "ErrServer")

		return
	}
	m := make(map[string]Leader)
	for _, list := range tables.Table {
		m[list[2]] = Leader{Organization: list[0], LeaderName: list[1]}
	}
	//fmt.Println(LeaderMap)
	LeaderMap = m
	//fmt.Println(LeaderMap)
	c.AsciiJSON(200, gin.H{
		"message": "success",
	})

}

func ErrResponse(c *gin.Context, msg string) {
	c.AsciiJSON(400, gin.H{
		"message": msg,
	})
}

func (u User) retire(account string) error {
	return dao.DB.Model(&User{}).Where("account = ", account).Error
}
