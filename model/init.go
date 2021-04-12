package model

import (
	"log"
	"nspyf/util"
	"os"
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

func LogInit(path string) {
	file, err := os.OpenFile(path,os.O_RDWR | os.O_CREATE | os.O_APPEND,666)
	if err != nil {
		panic(err)
	}

	ErrLog = log.New(file,"ERR",log.LstdFlags | log.Llongfile)
	return
}
