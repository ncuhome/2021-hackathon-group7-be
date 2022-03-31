package putable

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

var (
	jwtKey       = []byte("JwtKey")
	expireMinute = time.Duration(24)
)

type Claim struct {
	Account string
	jwt.StandardClaims
}

func GenerateToken(account string) string {

	var claims = Claim{
		Account: account,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),                               //立即生效
			ExpiresAt: time.Now().Add(expireMinute * time.Hour).Unix(), //失效时间
			Issuer:    "ncuQA-sever",                                   //签发者
		},
	}

	res := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	Token, err := res.SignedString(jwtKey)
	if err != nil {
		fmt.Println("err:", err)
		return ":"
	}

	//fmt.Println(Token)
	return "QAQ " + Token
}

func ParseToken(tokenString string) (claim *Claim) {
	token, err := jwt.ParseWithClaims(tokenString, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return string(jwtKey), nil
	})

	if err != nil {
		fmt.Println(err)
	}
	return token.Claims.(*Claim)
}

func Jwt(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.AsciiJSON(400, gin.H{
			"message": "请求参数错误",
		})
		c.Abort()
		return
	}
	parts := strings.SplitN(authHeader, " ", 2)
	//fmt.Println(parts)

	if !(len(parts) == 2 && parts[0] == "QAQ") {
		c.AsciiJSON(400, gin.H{
			"message": "错误token",
		})
		c.Abort()
		return
	}

	claim := ParseToken(parts[1])

	if time.Now().Unix() > claim.ExpiresAt {
		c.AsciiJSON(400, gin.H{
			"message": "token已过期,重新登录",
		})
		c.Abort()
		return
	}
	c.Next()

}
