package model

import (
	"log"
	"os"
	"tudo/util"
)

type JwtConfig struct {
	JwtTime int64 `json:"jwt_time"`
}

type EmailConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

var JwtConfigObj JwtConfig
var EmailConfigObj EmailConfig

var Jwt JWT
var Email GoMail
var ErrLog *log.Logger
var Admin map[string]string
var PictureObj *Picture
var OssObj *OssType

func JwtInit(path string) {
	if err := util.ReadJSON(path, &JwtConfigObj); err != nil {
		panic(err)
	}

	Jwt.Init(os.Getenv("JWT_SIGN_KEY"), JwtConfigObj.JwtTime)
	return
}

func EmailInit(path string) {
	if err := util.ReadJSON(path, &EmailConfigObj); err != nil {
		panic(err)
	}
	Email.Init(EmailConfigObj.Host, EmailConfigObj.Port, EmailConfigObj.Username, os.Getenv("EMAIL_PASSWORD"), EmailConfigObj.Name)
	return
}

func LogInit() {
	path := "./logs/err.txt"
	_, err := os.Stat("./logs")
	if err != nil {
		err := os.Mkdir("logs", os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 666)
	if err != nil {
		panic(err)
	}

	ErrLog = log.New(file, "ERR: ", log.LstdFlags|log.Llongfile)
	return
}

func AdminInit(path string) {
	if err := util.ReadJSON(path, &Admin); err != nil {
		panic(err)
	}
	return
}

func PictureInit(path string) {
	_, err := os.Stat(path)
	if err != nil {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	PictureObj.Path = path
	return
}

func OssInit() {
	OssObj = &OssType{
		Endpoint:        os.Getenv("OSS_ENDPOINT"),
		AccessKeyID:     os.Getenv("OSS_ACCESS_KEY_ID"),
		AccessKeySecret: os.Getenv("OSS_ACCESS_KEY_SECRET"),
		Bucket:          os.Getenv("OSS_BUCKET"),
	}
	return
}
