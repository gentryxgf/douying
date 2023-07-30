package middleware

import (
	"testing"
)

func TestJwt(t *testing.T) {
	var id uint = 3
	token, err := GenerateToken(id)
	t.Log(token, err)

	orgToken, err := ParseToken(token)
	t.Logf("%+v,%+v", orgToken, err)
}
