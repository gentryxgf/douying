package utils

import (
	"testing"
)

func TestMd5Encoder(t *testing.T) {
	data := "1234"
	str := Md5Encoder(data)
	t.Log(str)
}
