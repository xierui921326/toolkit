package utils

import (
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// InStringSlice 判断字符串是否在字符串切片中
//
//	@Description: 判断字符串是否在字符串切片中，如果在则返回true，否则返回false
//	@param target 要查找的字符串
//	@param arr 字符串切片
//	@return bool
func InStringSlice(target string, arr []string) bool {
	tmpList := make([]string, len(arr))
	// 目标的修改不会影响到原数组
	copy(tmpList, arr)
	// 对字符串切片进行排序
	sort.Strings(tmpList)
	index := sort.SearchStrings(tmpList, target)
	// 先判断 &&左侧的条件，如果不满足则结束此处判断，不会再进行右侧的判断
	if index < len(tmpList) && tmpList[index] == target {
		return true
	}
	return false
}

// StringToInt 将字符串转换为int类型
//
//	@Description: 将字符串转换为int类型
//	@param str 字符串
//	@return int
func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}

// StringToInt32 将字符串转换为int32类型
//
//	@Description: 将字符串转换为int32类型
//	@param str 字符串
//	@return int32
func StringToInt32(str string) int32 {
	i, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0
	}
	return int32(i)
}

// StringToInt64 将字符串转换为int64类型
//
//	@Description: 将字符串转换为int64类型
//	@param str 字符串
//	@return int64
func StringToInt64(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

// StringToFloat64 将字符串转换为float64类型
//
//	@Description: 将字符串转换为float64类型
//	@param str 字符串
//	@return float64
func StringToFloat64(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return f
}

// StringToFloat32 将字符串转换为float32类型
//
//	@Description: 将字符串转换为float32类型
//	@param str 字符串
//	@return float32
func StringToFloat32(str string) float32 {
	f, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0
	}
	return float32(f)
}

// InIntSlice 判断数字是否在数组内
//
//	@Description: 判断数字是否在数组内
//	@param target 要查找的数字
//	@param arr 数字数组
//	@return bool
func InIntSlice(target int, arr []int) bool {
	sort.Ints(arr)
	index := sort.SearchInts(arr, target)
	if index < len(arr) && arr[index] == target {
		return true
	}
	return false
}

// InInt64Slice 判断数字是否在数组内
//
//	@Description: 判断数字是否在数组内
//	@param target 要查找的数字
//	@param arr 数字数组
//	@return bool
func InInt64Slice(target int64, arr []int64) bool {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	index := sort.Search(len(arr), func(i int) bool { return arr[i] >= target })
	if index < len(arr) && arr[index] == target {
		return true
	}
	return false
}

// Int64ToString 将int64类型转换为字符串类型
//
//	@Description: 将int64类型转换为字符串类型
//	@param data int64类型数据
//	@return string
func Int64ToString(data int64) string {
	formatInt := strconv.FormatInt(data, 10)
	return formatInt
}

// Int32ToString 将int32类型转换为字符串类型
//
//	@Description: 将int32类型转换为字符串类型
//	@param data int32类型数据
//	@return string
func Int32ToString(data int32) string {
	formatInt := strconv.FormatInt(int64(data), 10)
	return formatInt
}

// IntToString 将int类型转换为字符串类型
//
//	@Description: 将int类型转换为字符串类型
//	@param data int类型数据
//	@return string
func IntToString(data int) string {
	formatInt := strconv.Itoa(data)
	return formatInt
}

// Uint64ToString 将uint64类型转换为字符串类型
//
//	@Description: 将uint64类型转换为字符串类型
//	@param data uint64类型数据
//	@return string
func Uint64ToString(data uint64) string {
	formatInt := strconv.FormatUint(data, 10)
	return formatInt
}

// Uint32ToString 将uint32类型转换为字符串类型
//
//	@Description: 将uint32类型转换为字符串类型
//	@param data uint32类型数据
//	@return string
func Uint32ToString(data uint32) string {
	formatInt := strconv.FormatUint(uint64(data), 10)
	return formatInt
}

// UintToString 将uint类型转换为字符串类型
//
//	@Description: 将uint类型转换为字符串类型
//	@param data uint类型数据
//	@return string
func UintToString(data uint) string {
	formatInt := strconv.FormatUint(uint64(data), 10)
	return formatInt
}

// Float64ToString 将float64类型转换为字符串类型
//
//	@Description: 将float64类型转换为字符串类型
//	@param data float64类型数据
//	@return string
func Float64ToString(data float64) string {
	formatFloat := strconv.FormatFloat(data, 'f', -1, 64)
	return formatFloat
}

// Float32ToString 将float32类型转换为字符串类型
//
//	@Description: 将float32类型转换为字符串类型
//	@param data float32类型数据
//	@return string
func Float32ToString(data float32) string {
	formatFloat := strconv.FormatFloat(float64(data), 'f', -1, 32)
	return formatFloat
}

// InterfaceIntoString 将interface{}类型转换为字符串类型
//
//	@Description: 将interface{}类型转换为字符串类型
//	@param data interface{}类型数据
//	@return string
func InterfaceIntoString(data interface{}) string {
	if data == nil {
		return ""
	}
	return string(data.([]byte))
}

// RemoveQuotes 删除字符串中单引号
//
//	@Description: 删除字符串中双引号
//	@param input 输入字符串
//	@return string 输出字符串
func RemoveQuotes(input string) string {
	return strings.ReplaceAll(input, "\"", "")
}

// IsUpperCase 如果在A-Z中的符文返回true
//
//	@Description: 如果在A-Z中的符文返回true
//	@param r 符文
//	@return bool
func IsUpperCase(r rune) bool {
	if r >= 'A' && r <= 'Z' {
		return true
	}
	return false
}

// IsLowerCase 如果a-z中的符文返回true
//
//	@Description: 如果a-z中的符文返回true
//	@param r 符文
//	@return bool
func IsLowerCase(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}
	return false
}

// ToSnakeCase
//
//	@Description: 通过将驼峰格式转换为蛇格式返回一个复制字符串
//	@param s 驼峰格式字符串
//	@return string 下划线字符串
func ToSnakeCase(s string) string {
	var out []rune
	for index, r := range s {
		if index == 0 {
			out = append(out, ToLowerCase(r))
			continue
		}

		if IsUpperCase(r) && index != 0 {
			if IsLowerCase(rune(s[index-1])) {
				out = append(out, '_', ToLowerCase(r))
				continue
			}
			if index < len(s)-1 && IsLowerCase(rune(s[index+1])) {
				out = append(out, '_', ToLowerCase(r))
				continue
			}
			out = append(out, ToLowerCase(r))
			continue
		}
		out = append(out, r)
	}
	return string(out)
}

// ToCamelCase
//
//	@Description: 通过将蛇形大小写转换为驼峰大小写返回一个复制字符串
//	@param s
//	@return string
func ToCamelCase(s string) string {
	s = ToLower(s)
	out := []rune{}
	for index, r := range s {
		if r == '_' {
			continue
		}
		if index == 0 {
			out = append(out, ToUpperCase(r))
			continue
		}

		if index > 0 && s[index-1] == '_' {
			out = append(out, ToUpperCase(r))
			continue
		}

		out = append(out, r)
	}
	return string(out)
}

// ToLowerCase 将符文转换为小写
//
//	@Description: 将符文转换为小写
//	@param r 符文
//	@return rune 小写符文
func ToLowerCase(r rune) rune {
	dx := 'A' - 'a'
	if IsUpperCase(r) {
		return r - dx
	}
	return r
}

// ToUpperCase
//
//	@Description: 将符文转换为大写
//	@param r 符文
//	@return rune 大写符文
func ToUpperCase(r rune) rune {
	dx := 'A' - 'a'
	if IsLowerCase(r) {
		return r + dx
	}
	return r
}

// ToLower 将字符串转为小写
//
//	@Description: 将字符串转为小写
//	@param s 字符串
//	@return string 小写字符串
func ToLower(s string) string {
	var out []rune
	for _, r := range s {
		out = append(out, ToLowerCase(r))
	}
	return string(out)
}

// ToUpper 将小写字母转为大写返回一个复制字符串
//
//	@Description: 将小写字母转为大写返回一个复制字符串
//	@param s 小写字母字符串
//	@return string 大写字母字符串
func ToUpper(s string) string {
	var out []rune
	for _, r := range s {
		out = append(out, ToUpperCase(r))
	}
	return string(out)
}

// UpperFirst 将第一个字母转换为大写
//
//	@Description: 将第一个字母转换为大写
//	@param s 字符串
//	@return string 转换后的字符串
func UpperFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return ToUpper(s[:1]) + s[1:]
}

// UnExport 将第一个字母转换为小写
//
//	@Description: 将第一个字母转换为小写
//	@param text 字符串
//	@return string 转换后的字符串
func UnExport(text string) string {
	var flag bool
	str := strings.Map(func(r rune) rune {
		if flag {
			return r
		}
		if unicode.IsLetter(r) {
			flag = true
			return unicode.ToLower(r)
		}
		return r
	}, text)
	return str
}

// DeleteSlice 删除指定元素
//
//	@Description: 删除指定元素
//	@param a 字符串切片
//	@param elem 要删除的元素
//	@return []string 新的字符串切片
func DeleteSlice(a []string, elem string) []string {
	j := 0
	for _, v := range a {
		if v != elem {
			a[j] = v
			j++
		}
	}
	return a[:j]
}

// NumberFormat 格式化数值
//
//	@Description: 格式化数值    1,234,567,898.55
//	@param str 字符串
//	@return string 格式化后的字符串
func NumberFormat(str string) string {
	length := len(str)
	if length < 4 {
		return str
	}
	arr := strings.Split(str, ".") //用小数点符号分割字符串,为数组接收
	length1 := len(arr[0])
	if length1 < 4 {
		return str
	}
	count := (length1 - 1) / 3
	for i := 0; i < count; i++ {
		arr[0] = arr[0][:length1-(i+1)*3] + "," + arr[0][length1-(i+1)*3:]
	}
	return strings.Join(arr, ".") //将一系列字符串连接为一个字符串，之间用sep来分隔。
}
