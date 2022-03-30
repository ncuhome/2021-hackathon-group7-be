package putable

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

var (
	jwtKey       = []byte("config.C.Token.JwtKey")
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
	//uid := c.Query("user_id")
	//fmt.Println(authHeader,uid)
	if authHeader == "" {
		//response.IllegalAccess(c)
		c.Abort()
		return
	}
	parts := strings.SplitN(authHeader, " ", 2)
	//parts := strings.Fields(authHeader)
	fmt.Println(parts)
	if parts[0] != "QAQ" || len(parts) != 2 {
		//response.WrongToken(c)
		c.Abort()
		return
	}
	claim := ParseToken(parts[1])

	var user User
	//database.DB.Where("id = ?", uid).Take(&user)
	if user.Account != claim.Account {
		//response.InvalidToken(c)
		c.Abort()
		return
	}
	if time.Now().Unix() > claim.ExpiresAt {
		//response.OverTimedToken(c)
		c.Abort()
		return
	}
	c.Next()

}
