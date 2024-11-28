package strutil

import (
	"strings"
)

// IsEmpty 判断字符串是否为空
func IsEmpty(str string) bool {
	return str == ""
}

// Length 返回字符串的长度
func Length(str string) int {
	return len(str)
}

// Concat 拼接多个字符串
func Concat(strs ...string) string {
	return strings.Join(strs, "")
}

// Replace 替换字符串中的子串
func Replace(str, old, new string, n int) string {
	return strings.Replace(str, old, new, n)
}

// TrimSpace 去除字符串首尾的空白字符
func TrimSpace(str string) string {
	return strings.TrimSpace(str)
}

// UpperCase 将字符串转换为大写
func UpperCase(str string) string {
	return strings.ToUpper(str)
}

// LowerCase 将字符串转换为小写
func LowerCase(str string) string {
	return strings.ToLower(str)
}
