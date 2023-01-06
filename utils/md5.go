package utils

import (
	md52 "crypto/md5"
	"encoding/hex"
	"io"
	"strings"
)

// Md5Encode 字符加密转小写
func Md5Encode(text string) string {
	md5 := md52.New()
	_, _ = io.WriteString(md5, text)
	return hex.EncodeToString(md5.Sum(nil))
}

// MD5Encode 字符串加密转大写
func MD5Encode(text string) string {
	md5 := Md5Encode(text)
	return strings.ToUpper(md5)
}

// MakeSaltPwd 盐值加密
func MakeSaltPwd(plainPwd, salt string) string {
	return Md5Encode(plainPwd + salt)
}

// ValidSaltPwd 解密 其实就是拿 当前加密值 与 库中的密码(加密串)做对比
func ValidSaltPwd(plainPwd, salt, saltPwd string) bool {
	return Md5Encode(plainPwd+salt) == saltPwd
}
