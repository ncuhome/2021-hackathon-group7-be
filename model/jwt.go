package model

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWT struct {
	signKey   string
	validTime int64
}

func (s *JWT) Init(signKey string, validTime int64) {
	s.signKey = signKey
	s.validTime = validTime
	return
}

func (s *JWT) GenerateToken(subject string, id string) (string, error) {
	now := time.Now().Unix()

	claims := jwt.StandardClaims{
		ExpiresAt: now + s.validTime,
		Id:        id,      //descp jwt状态 		 为loginStatus
		Subject:   subject, //descp token签发对象 为用户id gorm自动生成的
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(s.signKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

//验证签名并获取Claims
func (s *JWT) ParseToken(tokenString string) (*jwt.StandardClaims, error) {
	if tokenString == "" {
		return nil, errors.New("token is blank")
	}

	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.signKey), nil
	})

	if err != nil {
		return nil, err
	}

	if token == nil {
		return nil, errors.New("token解析为空")
	}

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token解析失败(UnknownError)")
}
