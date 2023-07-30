package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("jwtSecret")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// 颁发token
func GenerateToken(userid uint) (string, error) {
	expireAt := time.Now().Add(7 * 24 * time.Hour).Unix()

	claims := Claims{
		UserId: userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt,
			IssuedAt:  time.Now().Unix(),
			Issuer:    "gentry",
			Subject:   "douying",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// 解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
