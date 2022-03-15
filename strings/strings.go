package strings

import "unsafe"

// StringToBytes string 转 byte数组
func StringToBytes(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{str, len(str)},
	))
}

// BytesToString 字符数组转换为字符串，不需要创建新的内存
func BytesToString(byte []byte) string {
	return *(*string)(unsafe.Pointer(&byte))
}
