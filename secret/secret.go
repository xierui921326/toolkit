package secret

import (
	"crypto/sha1"
	"fmt"
	"github.com/xierui921326/toolkit/snow_node"
	"github.com/xierui921326/toolkit/utils"
	"io"
	"math"
	"strings"
	"time"
)

var tenToAny = map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f", 16: "g", 17: "h", 18: "i", 19: "j", 20: "k", 21: "l", 22: "m", 23: "n", 24: "o", 25: "p", 26: "q", 27: "r", 28: "s", 29: "t", 30: "u", 31: "v", 32: "w", 33: "x", 34: "y", 35: "z", 36: "A", 37: "B", 38: "C", 39: "D", 40: "E", 41: "F", 42: "G", 43: "H", 44: "I", 45: "J", 46: "K", 47: "L", 48: "M", 49: "N", 50: "O", 51: "P", 52: "Q", 53: "R", 54: "S", 55: "T", 56: "U", 57: "V", 58: "W", 59: "X", 60: "Y", 61: "Z"}

// decimalToAny
//
// @Description: 10进制转任意进制
// @param num: 10进制数
// @param n: 目标进制
// return: 目标进制数
func decimalToAny(num, n int) string {
	newNumStr := ""
	var remainder int
	var remainderString string
	for num != 0 {
		remainder = num % n
		remainderString = tenToAny[remainder]
		newNumStr = remainderString + newNumStr
		num = num / n
	}
	return newNumStr
}

// findKey
//
// @Description: 找出value对应的key
// @param in: 目标字符串
// return: key
func findKey(in string) int {
	result := -1
	for k, v := range tenToAny {
		if in == v {
			result = k
		}
	}
	return result
}

// anyToDecimal
//
// @Description: 任意进制转10进制
// @param num: 任意进制数
// @param n: 目标进制
// return: 10进制数
func anyToDecimal(num string, n int) int {
	var newNum float64
	newNum = 0.0
	nNum := len(strings.Split(num, "")) - 1
	for _, value := range strings.Split(num, "") {
		tmp := float64(findKey(value))
		if tmp != -1 {
			newNum = newNum + tmp*math.Pow(float64(n), float64(nNum))
			nNum = nNum - 1
		} else {
			break
		}
	}
	return int(newNum)
}

// GetAppId
//
// @Description: 生成APP ID
// return: APP ID
func GetAppId() string {
	id := snow_node.GetID()
	return decimalToAny(utils.StringIntoInt(id), 62)
}

// GetAppKey
//
// @Description: 生成APP Key
// return: APP Key
func GetAppKey() string {
	id := snow_node.GetID()
	return decimalToAny(utils.StringIntoInt(id), 32)
}

// GetAppSecret
//
// @Description: 生成App Secret 算法： sha1(appId+timestamp) 生成AppSecret
// @param appId: APP ID
// return: APP Secret
func GetAppSecret(appId string) string {
	//定义Buffer类型
	var bt strings.Builder
	bt.WriteString(appId)
	currentTime := time.Now().UnixMilli()
	bt.WriteString(utils.Int64IntoString(currentTime))

	//对字符串进行SHA1哈希
	t := sha1.New()
	io.WriteString(t, bt.String())
	return fmt.Sprintf("%x", t.Sum(nil))
}
