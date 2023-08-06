package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(src []byte) string {
	m := md5.New()
	m.Write(src)
	return hex.EncodeToString(m.Sum([]byte("")))
}
