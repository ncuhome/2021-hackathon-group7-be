package service

import (
	"encoding/hex"
	"fmt"
	"github.com/tidwall/gjson"
	"strconv"
	"time"
	"tudo/util"
)

var DocSource = &DocBaseData{
	Url: "https://docs.qq.com/dop-api/opendoc?tab=BB08J2&id=DS3B5S2h6bktxZHJr&outformat=1&normal=1",
}

type Leader struct {
	Organization string
	LeaderName   string
}

var LeaderMap = map[string]Leader{} // 电话：leader

type DocBaseData struct {
	Url          string
	JsonCacheMD5 string
}

func GetCell(tableData *gjson.Result, x int64, y int64) string {
	place := strconv.FormatInt(x*26+y, 10)
	arr := tableData.Map()[place].Map()["2"].Array()
	if len(arr) < 2 {
		return ""
	}
	data := arr[1].String()
	return data
}

func JsonToEmailMap(table *gjson.Result) uint {
	arr := table.Map()["c"].Array()
	if len(arr) < 2 {
		return ErrorServer
	}
	tableData := arr[1]

	LeaderMap = map[string]Leader{}
	for i := int64(1); ; i++ {
		organization := GetCell(&tableData, i, 0)
		if organization == "" {
			break
		}

		leaderName := GetCell(&tableData, i, 1)
		phone := GetCell(&tableData, i, 2)

		if phone == "" {
			continue
		}
		LeaderMap[phone] = Leader{
			Organization: organization,
			LeaderName:   leaderName,
		}
	}

	fmt.Println("ToMap OK")
	return SuccessCode
}

func GetTableA1(table *gjson.Result) string {
	arr := table.Map()["c"].Array()
	if len(arr) < 2 {
		return ""
	} else {
		arr = arr[1].Map()["0"].Map()["2"].Array()
		if len(arr) < 2 {
			return ""
		} else if arr[1].String() == "" {
			return ""
		}
	}
	return arr[1].String()
}

func GetDocs(docBaseData *DocBaseData) (*gjson.Result, uint) {
	option := &util.HttpOption{
		Url:    docBaseData.Url,
		Method: "GET",
	}

	data, err := util.HttpReq(option)
	if err != nil {
		fmt.Println(err)
		return nil, ErrorServer
	}

	// ? 不知道为啥json可能不一样，根据能否获取第一行第一列来选
	table := gjson.GetBytes(data, "clientVars.collab_client_vars.initialAttributedText.text.0.1.0")
	A1 := GetTableA1(&table)
	for i := 0; i < 10; i++ {
		table = gjson.GetBytes(data, "clientVars.collab_client_vars.initialAttributedText.text.0."+strconv.Itoa(i)+".0")
		A1 = GetTableA1(&table)
		if A1 != "" {
			break
		}
	}

	if A1 == "" {
		fmt.Println("json格式有问题")
		return nil, ErrorServer
	}

	// 暂停同步
	// fmt.Println(A1[0:3])
	if A1[0:3] == "QAQ" {
		return nil, SuccessCode
	}

	md5 := hex.EncodeToString(util.MD5([]byte(table.String())))
	if docBaseData.JsonCacheMD5 == md5 {
		return nil, SuccessCode
	}
	docBaseData.JsonCacheMD5 = md5
	return &table, SuccessCode
}

func TencentDocToMap(docBaseData *DocBaseData) uint {
	table, code := GetDocs(docBaseData)

	if code != SuccessCode {
		return code
	}
	if table == nil {
		return SuccessCode
	}

	code = JsonToEmailMap(table)
	if code != SuccessCode {
		return code
	}

	return SuccessCode
}

func SyncTencentDoc() {
	go func() {
		for {
			TencentDocToMap(DocSource)
			time.Sleep(time.Hour)
			// time.Sleep(time.Second * 5)
		}
	}()
}

func TestTencentDoc() {
	LeaderMap["15797702607"] = Leader{
		Organization: "前端测试",
		LeaderName:   "黄",
	}
	LeaderMap["15107076230"] = Leader{
		Organization: "后端测试",
		LeaderName:   "彭",
	}
}
