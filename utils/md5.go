package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5加密,返回大写字符串
func Md5Encoder(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
