package encrypt

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"github.com/xierui921326/toolkit/utils"
)

// Sha256 实现sha256加密算法
type Sha256 struct {
	// Key 加密密钥
	Key string
}

// NewSha256 创建一个Sha256对象
//
// @Description: 创建一个Sha256对象
// @param key 加密密钥
// @return *Sha256 Sha256对象
func NewSha256(key string) *Sha256 {
	return &Sha256{
		Key: key,
	}
}

// 加密
func (s *Sha256) encrypt(message string) []byte {
	key := []byte(s.Key)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))

	return h.Sum(nil)
}

// ToHex 将加密后的二进制转16进制字符串
//
// @Description: 由于sha256加密后的结果是16进制字符串，所以这里直接转base64时，需要将结果转为大写字母
// @param message 待加密的字符串
// @return string 加密后的16进制字符串
func (s *Sha256) ToHex(message string) string {
	return hex.EncodeToString(s.encrypt(message))
}

// ToStdBase64 将加密后的二进制转标准base64
//
// @Description: 由于sha256加密后的结果是16进制字符串，所以这里直接转base64时，需要将结果转为大写字母
// @param message 待加密的字符串
// @return string 加密后的base64字符串
func (s *Sha256) ToStdBase64(message string) string {
	return base64.StdEncoding.EncodeToString(s.encrypt(message))
}

// ToUrlBase64 将加密后的二进制转Url base64
//
// @Description: 由于sha256加密后的结果是16进制字符串，所以这里直接转base64时，需要将结果转为大写字母
// @param message 待加密的字符串
// @return string 加密后的base64字符串
func (s *Sha256) ToUrlBase64(message string) string {
	return base64.URLEncoding.EncodeToString(s.encrypt(message))
}

// Sha1 实现sha1加密算法
type Sha1 struct {
	// Key 加密密钥
	Key string
}

// NewSha1 创建一个Sha1对象
//
// @Description: 创建一个Sha1对象
// @param key 加密密钥
// @return *Sha1 Sha1对象
func NewSha1(key string) *Sha1 {
	return &Sha1{
		Key: key,
	}
}

// 加密
func (s *Sha1) encrypt(message string) []byte {
	key := []byte(s.Key)
	h := hmac.New(sha1.New, key)
	h.Write([]byte(message))

	return h.Sum(nil)
}

// ToHex 将加密后的二进制转16进制字符串
//
// @Description: 由于sha1加密后的结果是16进制字符串，所以这里直接转base64时，需要将结果转为大写字母
// @param message 待加密的字符串
// @return string 加密后的16进制字符串
func (s *Sha1) ToHex(message string) string {
	return hex.EncodeToString(s.encrypt(message))
}

// ToStdBase64 将加密后的二进制转标准base64
//
// @Description: 由于sha1加密后的结果是16进制字符串，所以这里直接转base64时，需要将结果转为大写字母
// @param message 待加密的字符串
// @return string 加密后的base64字符串
func (s *Sha1) ToStdBase64(message string) string {
	return base64.StdEncoding.EncodeToString([]byte(utils.ToUpper(message)))
}

// ToUrlBase64 将加密后的二进制转Url base64
//
// @Description: 由于sha1加密后的结果是16进制字符串，所以这里直接转base64时，需要将结果转为大写字母
// @param message 待加密的字符串
// @return string 加密后的base64字符串
func (s *Sha1) ToUrlBase64(message string) string {
	return base64.URLEncoding.EncodeToString(s.encrypt(message))
}
