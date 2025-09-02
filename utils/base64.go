package utils

import "encoding/base64"

// Base64Encode base64加密
//
// @Description: base64加密
// @param decoded 待加密数据
// @return 加密后数据
func Base64Encode(decoded []byte) string {
	return base64.StdEncoding.EncodeToString(decoded)
}

// Base64Decode base64解密
//
// @Description: base64解密
// @param encoded 待解密数据
// @return 解密后数据
// @return error 错误信息
func Base64Decode(encoded string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encoded)
}
