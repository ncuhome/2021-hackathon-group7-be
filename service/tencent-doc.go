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

var EmailMap = map[string]int{}

type DocBaseData struct {
	Url          string
	JsonCacheMD5 string
}

func JsonToEmailMap(table *gjson.Result) uint {
	arr := table.Map()["c"].Array()
	if len(arr) < 2 {
		return ServerError
	}
	tableData := arr[1]

	EmailMap = map[string]int{}
	for i := int64(1); ; i++ {
		place := strconv.FormatInt(i*26+0, 10)
		arr = tableData.Map()[place].Map()["2"].Array()
		if len(arr) < 2 {
			break
		}
		email := arr[1].String()
		if email == "" {
			break
		}

		EmailMap[email] = 1
	}

	fmt.Println("ToEmailMap OK")
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
		return nil, ServerError
	}

	// ? 不知道为啥json可能不一样，根据能否获取第一行第一列来选
	table := gjson.GetBytes(data, "clientVars.collab_client_vars.initialAttributedText.text.0.1.0")
	A1 := GetTableA1(&table)
	if A1 == "" {
		table = gjson.GetBytes(data, "clientVars.collab_client_vars.initialAttributedText.text.0.2.0")
		A1 = GetTableA1(&table)
	}

	if A1 == "" {
		fmt.Println("json格式有问题")
		return nil, ServerError
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

func TencentDocToES(docBaseData *DocBaseData) uint {
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
		for ; ; {
			TencentDocToES(DocSource)
			time.Sleep(time.Second * 5)
			// time.Sleep(time.Minute * 5)
		}
	}()
}
