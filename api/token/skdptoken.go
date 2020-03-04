package token

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	SecretKey = "4color.cn"
)

func GetJwt() (jwtJson string) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	tokenString, _ := token.SignedString([]byte(SecretKey))

	jwtJson = tokenString

	return
}
