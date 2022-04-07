package model

import (
	"log"
	"os"
	"tudo/util"
)

type JwtConfig struct {
	JwtTime int64 `json:"jwt_time"`
}

var JwtConfigObj JwtConfig

var Jwt JWT
var ErrLog *log.Logger
var OssObj *OssType
var OssBaseUrl string

func JwtInit(path string) {
	if err := util.ReadJSON(path, &JwtConfigObj); err != nil {
		panic(err)
	}

	Jwt.Init(os.Getenv("JWT_SIGN_KEY"), JwtConfigObj.JwtTime)
	return
}

func LogInit() {
	path := "./logs/err.txt"
	_, err := os.Stat("./logs")
	if err != nil {
		err = os.Mkdir("logs", os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 666)
	if err != nil {
		panic(err)
	}

	ErrLog = log.New(file, "ERR: ", log.LstdFlags|log.Llongfile) //ErrLog会收集err,
	return
}

func OssInit() {
	OssObj = &OssType{
		Endpoint:        os.Getenv("OSS_ENDPOINT"),
		AccessKeyID:     os.Getenv("OSS_ACCESS_KEY_ID"),
		AccessKeySecret: os.Getenv("OSS_ACCESS_KEY_SECRET"),
		Bucket:          os.Getenv("OSS_BUCKET"),
	}
	OssBaseUrl = "https://" + OssObj.Bucket + "." + OssObj.Endpoint + "/" //descp oss的基础url (在后面加上文件名即可)
	return
}
