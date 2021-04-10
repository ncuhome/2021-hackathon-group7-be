package model

import (
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
