package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

const (
	BlockSize = 16
	CBC       = "CBC"
	ECB       = "ECB"
)

// unPaddingFuncMap 去填充函数映射
var unPaddingFuncMap = map[AesType]UnPaddingFunc{
	AesTypeCBCZeroPadding:  zeroUnPadding,
	AesTypeCBCPKCS7Padding: pKCS7UnPadding,
}

// paddingFuncMap 填充函数映射
var paddingFuncMap = map[AesType]PaddingFunc{
	AesTypeCBCPKCS7Padding: pKCS7Padding,
	AesTypeCBCZeroPadding:  zeroPadding,
}

const (
	// AesTypeCBCZeroPadding 0填充
	AesTypeCBCZeroPadding AesType = iota
	// AesTypeCBCPKCS7Padding PKCS7填充
	AesTypeCBCPKCS7Padding
)

// PaddingFunc 填充函数
//
// @Description: 填充函数
// @param data 需要填充的数据
// @param blockSize 块大小
// @return []byte 填充后的数据
type PaddingFunc func([]byte, int) []byte

// UnPaddingFunc 去填充函数
//
// @Description: 去填充函数
// @param data 需要去填充的数据
// @return []byte 去填充后的数据
type UnPaddingFunc func([]byte) []byte

// AesType 加密类型
type AesType int

type AesCbc struct {
	// AesType 填充类型
	AesType AesType `json:"aes_type"`
	// Key 加密密钥
	Key string `json:"key"`
	// Iv 初始向量
	Iv string `json:"iv"`
}

type AesEcb struct {
	// AesType 填充类型
	AesType AesType `json:"aes_type"`
	// Key 加密密钥
	Key string `json:"key"`
}

func NewAesCbc(aesType AesType, key, iv string) *AesCbc {
	return &AesCbc{
		AesType: aesType,
		Key:     key,
		Iv:      iv,
	}
}

func NewAesEcb(aesType AesType, key string) *AesEcb {
	return &AesEcb{
		AesType: aesType,
		Key:     key,
	}
}

// AesCbcDecrypt
//
//	@Description: aes-cbc解密
//	@param ciphertext 需要解密的文本
//	@return []byte 解密后的数据
//	@return error
func (a *AesCbc) AesCbcDecrypt(ciphertext string) ([]byte, error) {
	unPaddingFunc, ok := unPaddingFuncMap[a.AesType]
	if !ok {
		return nil, errors.New("unsupported Aes type")
	}
	return decrypt(unPaddingFunc, a.Key, a.Iv, CBC, ciphertext)
}

// AesCbcEncrypt
//
//	@Description: aes-cbc加密
//	@param aesType 填充类型
//	@param key 加密密钥
//	@param srcData 需要加密的文本数据
//	@return string 加密后的数据
//	@return error
func (a *AesCbc) AesCbcEncrypt(srcData []byte) (string, error) {
	paddingFn, ok := paddingFuncMap[a.AesType]
	if !ok {
		return "", errors.New("unsupported aes type")
	}
	return encrypt(paddingFn, a.Key, a.Iv, CBC, srcData)
}

func (a *AesEcb) AesEcbDecrypt(ciphertext string) ([]byte, error) {
	unPaddingFunc, ok := unPaddingFuncMap[a.AesType]
	if !ok {
		return nil, errors.New("unsupported Aes type")
	}
	return decrypt(unPaddingFunc, a.Key, "", ECB, ciphertext)
}

// AesEcbEncrypt
//
//	@Description: aes-ecb加密
//	@param aesType 填充类型
//	@param key 加密密钥
//	@param srcData 需要加密的文本数据
//	@return string 加密后的数据
//	@return error
func (a *AesEcb) AesEcbEncrypt(srcData []byte) (string, error) {
	paddingFn, ok := paddingFuncMap[a.AesType]
	if !ok {
		return "", errors.New("unsupported aes type")
	}
	return encrypt(paddingFn, a.Key, "", ECB, srcData)
}

/**
 * aes-ecb解密
 */
func decrypt(unPaddingFn UnPaddingFunc, key, iv, mode string, ciphertext string) ([]byte, error) {
	cKey, err := aes.NewCipher([]byte(key))
	if nil != err {
		return nil, err
	}
	var decrypter cipher.BlockMode
	switch mode {
	case CBC:
		decrypter = cipher.NewCBCDecrypter(cKey, []byte(iv))
	case ECB:
		decrypter = newECBDecrypter(cKey)
	}
	base64In, _ := base64.StdEncoding.DecodeString(ciphertext)
	in := make([]byte, len(base64In))
	decrypter.CryptBlocks(in, base64In)
	return unPaddingFn(in), nil
}

/**
 * aes-ecb加密
 */
func encrypt(paddingFn PaddingFunc, key, iv, mode string, srcData []byte) (string, error) {
	cKey, err := aes.NewCipher([]byte(key))
	if nil != err {
		return "", err
	}
	var encrypter cipher.BlockMode
	switch mode {
	case CBC:
		encrypter = cipher.NewCBCEncrypter(cKey, []byte(iv))
	case ECB:
		encrypter = newECBEncrypter(cKey)
	}

	srcData = paddingFn(srcData, BlockSize)
	out := make([]byte, len(srcData))
	encrypter.CryptBlocks(out, srcData)
	return base64.StdEncoding.EncodeToString(out), nil
}

/**
 * PKCS7补码
 */
func pKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

/**
 * PKCS7去码
 */
func pKCS7UnPadding(data []byte) []byte {
	length := len(data)
	// 去掉最后一个字节 unPadding 次
	unPadding := int(data[length-1])
	return data[:(length - unPadding)]
}

/**
 * 零填充
 */
func zeroPadding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(0)}, padding)
	return append(data, padText...)
}

/**
 * 零去码
 */
func zeroUnPadding(origData []byte) []byte {
	length := len(origData)
	i := 0
	for ; 0 == int(origData[length-1-i]); i++ {
	}
	return origData[:(length - i)]
}

/**
 * ecb模式加密
 */
type ecb struct {
	b         cipher.Block
	blockSize int
}

/**
 * 创建一个ecb模式的加密器
 * @param b 加密器使用的块
 */
func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

// ecb模式加密器
type ecbEncrypter ecb

/**
 * 创建一个ecb模式的加密器
 * @param b 加密器使用的块
 */
func newECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}

// BlockSize 返回块的大小
//
// @Description: 返回块的大小
// @return int
func (x *ecbEncrypter) BlockSize() int { return x.blockSize }

// CryptBlocks 加密块
//
// @Description: 加密块
// @param dst 加密后的数据
// @param src 加密前的数据
func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

// ecb模式解密器
type ecbDecrypter ecb

// 创建一个ecb模式的解密器
//
// @Description: 创建一个ecb模式的解密器
// @param b 解密器使用的块
// @return cipher.BlockMode
func newECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}

// BlockSize 返回块的大小
//
// @Description: 返回块的大小
// @return int
func (x *ecbDecrypter) BlockSize() int { return x.blockSize }

// CryptBlocks 解密块
//
// @Description: 解密块
// @param dst 解密后的数据
// @param src 解密前的数据
func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

// AesGcm aes-gcm加密解密
type AesGcm struct {
	// Key 加密密钥
	Key []byte
}

// Encrypt AES GCM模式加密
//
// @Description: AES GCM模式加密
// @param origData 原始数据
// @return string 加密后的数据
// @return error
func (a *AesGcm) Encrypt(origData string) (string, error) {
	if origData == "" {
		return "", errors.New("加密数据不能为空")
	}
	block, err := aes.NewCipher(a.Key)
	if err != nil {
		return "", err
	}

	// 创建一个GCM模式的加密器
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// 创建一个随机nonce
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// 加密数据
	out := aesGCM.Seal(nil, nonce, []byte(origData), nil)
	return base64.StdEncoding.EncodeToString(append(nonce, out...)), nil
}

// Decrypt AES GCM模式解密
//
// @Description: AES GCM模式解密
// @param crypted 加密后的数据
// @return string 解密后的数据
// @return error
func (a *AesGcm) Decrypt(crypted []byte) (string, error) {
	if len(a.Key) != 16 {
		return "", errors.New("key must be 16 bytes for AES-128")
	}

	block, err := aes.NewCipher(a.Key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(crypted) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce := crypted[:nonceSize]
	ciphertext := crypted[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
