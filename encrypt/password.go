package encrypt

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// Encrypt 加密密码
//
// @Description: 使用bcrypt算法加密密码
// @param password 密码明文
// @return string 密码密文
func Encrypt(password string) string {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// never runs here
		log.Printf("[NRH] encrypt failed error: %v", err)
		return ""
	}

	return string(bs)
}

// Compare 密码校验
//
// @Description: 使用bcrypt算法校验密码
// @param hashedPassword 密码密文
// @param password 密码明文
// @return bool 校验结果
func Compare(hashedPassword, password string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return false
	}

	if err != nil {
		// never runs here
		log.Printf("[NRH] compare failed error: %v", err)
		return false
	}

	return true
}
