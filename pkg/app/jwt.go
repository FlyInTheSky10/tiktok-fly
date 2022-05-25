package app

import (
	"github.com/FlyInThesky10/TikTok-Fly/global"
	"github.com/FlyInThesky10/TikTok-Fly/pkg/errcode"
	"github.com/golang-jwt/jwt"
	"time"
)

type Claims struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateJWTToken(username, password string) (string, int64, error) {
	claims := &Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(global.JWTSetting.Expire).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(global.JWTSetting.Secret))
	if err != nil {
		return t, 0, errcode.UnauthorizedTokenGenerate
	}
	return t, claims.ExpiresAt, nil
}
func ParseJWTToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.JWTSetting.Secret), nil
	})
	if err != nil {
		return nil, errcode.UnauthorizedTokenError
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	if !token.Valid {
		return nil, errcode.UnauthorizedTokenTimeout
	}
	return nil, errcode.UnauthorizedTokenError
}
