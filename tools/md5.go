package tools

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

//Sign 签名
func Sign(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	cipherStr := md5Ctx.Sum(nil)
	md5str := hex.EncodeToString(cipherStr)

	return strings.ToUpper(md5str)
}
