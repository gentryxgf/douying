package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go/v4"
	"go.uber.org/zap"
	"mini-tiktok/common/global"
	"time"
)

var MySecret []byte

type PayLoad struct {
	Username string `json:"username"`
	UserID   int    `json:"userID"`
}

type UserClaim struct {
	PayLoad
	jwt.StandardClaims
}

func GenToken(user PayLoad) (string, error) {
	MySecret = []byte(global.Config.JwtConf.Secret)
	claim := UserClaim{
		PayLoad: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(global.Config.JwtConf.Expires))),
			Issuer:    global.Config.JwtConf.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
}

func ParseToken(tokenStr string) (*UserClaim, error) {
	MySecret = []byte(global.Config.JwtConf.Secret)
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		global.Log.Error("TOKEN PARSE ERROR", zap.Error(err))
		return nil, err
	}
	if token != nil {
		if claim, ok := token.Claims.(*UserClaim); ok && token.Valid {
			return claim, nil
		}
	}
	return nil, errors.New("INVALID TOKEN")
}
