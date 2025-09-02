package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

// HmacSha256 计算HmacSha256
//
// @Description: 计算HmacSha256
// @param key 加密所使用的key
// @param data 加密的内容
// @return []byte 加密后的二进制
func HmacSha256(key string, data string) []byte {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))

	return mac.Sum(nil)
}

// HmacSha256ToHex 将加密后的二进制转16进制字符串
//
// @Description: 将加密后的二进制转16进制字符串
// @param key 加密所使用的key
// @param data 加密的内容
// @return string 加密后的16进制字符串
func HmacSha256ToHex(key string, data string) string {
	return hex.EncodeToString(HmacSha256(key, data))
}

// HmacSha256ToBase64 将加密后的二进制转Base64字符串
//
// @Description: 将加密后的二进制转Base64字符串
// @param key 加密所使用的key
// @param data 加密的内容
// @return string 加密后的Base64字符串
func HmacSha256ToBase64(key string, data string) string {
	return base64.URLEncoding.EncodeToString(HmacSha256(key, data))
}

// HmacSha1ToString 计算HmacSha1
//
// @Description: 计算HmacSha1
// @param keyStr 加密所使用的key
// @param data 加密的内容
// @return string 加密后的字符串
func HmacSha1ToString(keyStr, data string) string {
	// Crypto by HMAC-SHA1
	key := []byte(keyStr)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(data))

	//进行base64编码
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

// HmacSha1 计算HmacSha1
//
// @Description: 计算HmacSha1
// @param keyStr 加密所使用的key
// @param data 加密的内容
// @return []byte 加密后的二进制
func HmacSha1(keyStr, data string) []byte {
	// Crypto by HMAC-SHA1
	key := []byte(keyStr)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(data))

	return mac.Sum(nil)
}

// HmacSha1ToHex 将加密后的二进制转16进制字符串
//
// @Description: 将加密后的二进制转16进制字符串
// @param key 加密所使用的key
// @param data 加密的内容
// @return string 加密后的16进制字符串
func HmacSha1ToHex(key string, data string) string {
	return hex.EncodeToString(HmacSha1(key, data))
}
