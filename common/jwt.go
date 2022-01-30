package common
/*
jwt全称：JSON Web Tokens
 */


import (
	"github.com/dgrijalva/jwt-go"
	"learn/ginEssential/models"
	"time"
)


var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	//引用结构体
	jwt.StandardClaims
}

func ReleaseToken(user models.User) (string, error) {
	//Payload, 存在信息
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "oceanlearn.tech",
			Subject:   "user token",
		},
	}

	//加密签名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//获取token
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}


func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims:= &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}