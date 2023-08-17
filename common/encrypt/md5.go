package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

var salt = "douyin"

func Md5(src []byte) string {
	m := md5.New()
	m.Write(src)
	return hex.EncodeToString(m.Sum([]byte("")))
}

func Md5Encode(data string) string {
	m := md5.New()
	m.Write([]byte(data))
	return hex.EncodeToString(m.Sum(nil))
}

func GetPwd(plainpwd string) string {
	return Md5Encode(plainpwd + salt)
}

func ComparePwd(plainpwd, password string) bool {
	md := GetPwd(plainpwd)
	return md == password
}
