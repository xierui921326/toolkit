package utils

import (
	"archive/zip"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
)

// SortedQueryString 数据字典排序
//
// @description 排序请求参数字典，并返回排序后的请求参数字符串(不包含签名参数)
// @params params 请求参数字典
// @return string 排序后的请求参数字符串
func SortedQueryString(params map[string]string) string {
	keys := make([]string, 0, len(params))
	for k, _ := range params {
		if strings.EqualFold(k, "signature") {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var sortedQuery strings.Builder
	for i, each := range keys {
		if i > 0 {
			sortedQuery.WriteByte('&')
		}
		sortedQuery.WriteString(strings.ReplaceAll(url.QueryEscape(each),
			"+", "%20"))
		sortedQuery.WriteByte('=')
		sortedQuery.WriteString(strings.ReplaceAll(url.QueryEscape(params[each]), "+", "%20"))
	}
	return sortedQuery.String()
}

// CheckPort 检查端口是否存在
//
// @description 检查端口是否存在，返回 true 存在，false 不存在
// @params port 端口号
// @return bool 端口是否存在
func CheckPort(port int) bool {
	checkStatement := fmt.Sprintf(`netstat -anp | grep -q %d ; echo $?`, port)
	output, err := exec.Command("sh", "-c", checkStatement).CombinedOutput()
	if err != nil {
		return false
	}
	// log.println(output, string(output)) ==> [48 10] 0 或 [49 10] 1
	result, err := strconv.Atoi(strings.TrimSuffix(string(output), "\n"))
	if err != nil {
		return false
	}
	if result == 0 {
		return true
	}

	return false
}

// PathExists 判断文件或目录是否存在
//
// @description 判断文件或目录是否存在，返回 true 存在，false 不存在
// @params path 文件或目录路径
// @return bool 文件或目录是否存在
// @return error 错误信息
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// GetAllFile 获取指定目录下所有文件名
//
// @description 获取指定目录下所有文件名，返回文件名列表
// @params pathname 指定目录
// @return []string 文件名列表
// @return error 错误信息
func GetAllFile(pathname string, s []string) ([]string, error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		return s, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := pathname + "/" + fi.Name()
			s, err = GetAllFile(fullDir, s)
			if err != nil {
				return s, err
			}
		} else {
			fullName := fi.Name()
			s = append(s, fullName)
		}
	}
	return s, nil
}

// CreateDir 创建目录
//
// @description 创建目录，返回目录路径
// @params basePath 基础路径
// @params folderName 目录名称
// @return string 目录路径
func CreateDir(basePath, folderName string) (dirPath string) {
	folderPath := filepath.Join(basePath, folderName)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步
		// 先创建文件夹
		os.Mkdir(folderPath, 0777)
		// 再修改权限
		os.Chmod(folderPath, 0777)
	}
	return folderPath
}

// Zip 压缩文件
//
// @description 压缩文件，返回压缩文件大小
// @params srcFile 源文件路径
// @params destZip 压缩文件路径
// @return int64 压缩文件大小
func Zip(srcFile string, destZip string) int64 {
	// 预防：旧文件无法覆盖
	os.RemoveAll(destZip)

	// 创建：zip文件
	zipFile, _ := os.Create(destZip)
	defer zipFile.Close()

	// 打开：zip文件
	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	// 遍历路径信息
	filepath.Walk(srcFile, func(path string, info os.FileInfo, _ error) error {
		// 如果是源路径，提前进行下一个遍历
		if path == srcFile {
			return nil
		}

		// 获取：文件头信息
		header, _ := zip.FileInfoHeader(info)
		header.Name = strings.TrimPrefix(path, srcFile+`/`)

		// 判断：文件是不是文件夹
		if info.IsDir() {
			header.Name += `/`
		} else {
			// 设置：zip的文件压缩算法
			header.Method = zip.Deflate
		}

		// 创建：压缩包头部信息
		writer, _ := archive.CreateHeader(header)
		if !info.IsDir() {
			file, _ := os.Open(path)
			defer file.Close()
			io.Copy(writer, file)
		}
		return nil
	})

	fi, _ := zipFile.Stat()
	fileSize := fi.Size()

	return fileSize
}

// Interface2Type 转换接口类型
//
// @description 转换接口类型，返回转换后的类型值
// @params i 接口
// @return string 转换后的类型值
func Interface2Type(i interface{}) string {
	value := ""
	switch i.(type) {
	case string:
		value = i.(string)
		break
	case int:
		value = fmt.Sprintf("%d", i.(int))
		break
	case int8:
		value = fmt.Sprintf("%d", i.(int8))
		break
	case int16:
		value = fmt.Sprintf("%d", i.(int16))
		break
	case int32:
		value = fmt.Sprintf("%d", i.(int32))
		break
	case int64:
		value = fmt.Sprintf("%d", i.(int64))
		break
	case float64:
		value = strconv.FormatFloat(i.(float64), 'f', 3, 64)
		// 去除result末尾的0
		for strings.HasSuffix(value, "0") {
			value = strings.TrimSuffix(value, "0")
		}
		if strings.HasSuffix(value, ".") {
			value = strings.TrimSuffix(value, ".")
		}
		break
	default:
		value = fmt.Sprintf("%v", i)
		break
	}

	return value
}

// Decimal 保存指定小数点位数的浮点数
//
// @description 保存指定小数点位数的浮点数，返回保存后的浮点数
// @params value 浮点数
// @params pre 小数点位数
// @return float64 保存后的浮点数
func Decimal(value float64, pre int) float64 {
	v := strconv.FormatFloat(value, 'f', pre, 64)
	// 去除result末尾的0
	for strings.HasSuffix(v, "0") {
		v = strings.TrimSuffix(v, "0")
	}
	if strings.HasSuffix(v, ".") {
		v = strings.TrimSuffix(v, ".")
	}
	val, _ := strconv.ParseFloat(v, 64)
	return val
}

// Md5 计算字符串的MD5值
//
// @description 计算字符串的MD5值，返回MD5值
// @params str 字符串
// @return string MD5值
func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

// StructToMap 将结构体转换为 map[string]interface{}
//
// @description 将结构体转换为 map[string]interface{}，返回转换后的 map
// @params obj 结构体
// @params isSnake 是否将字段名转换为下划线格式
// @return map[string]interface{} 转换后的 map
func StructToMap(obj interface{}, isSnake bool) map[string]interface{} {
	// 创建一个空的 map
	result := make(map[string]interface{})

	// 检查是否为 nil
	if obj == nil {
		return result
	}

	// 获取结构体的反射值
	value := reflect.ValueOf(obj)

	// 检查是否是指针，如果是，我们获取元素的值
	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return result
		}
		value = value.Elem()
	}

	// 检查是否是结构体
	if value.Kind() != reflect.Struct {
		return result
	}

	// 获取结构体的类型
	t := value.Type()

	// 遍历结构体的字段
	for i := 0; i < value.NumField(); i++ {
		// 获取字段名
		field := t.Field(i)
		// 获取字段的值
		fieldValue := value.Field(i)

		// 跳过未导出字段(安全检查：确保字段可导出)
		if !fieldValue.CanInterface() {
			continue
		}

		// 如果字段名为空，跳过
		if field.Name == "" {
			continue
		}

		fieldName := field.Name
		if isSnake {
			// 将字段名从驼峰格式转换为下划线格式
			fieldName = strcase.ToSnake(field.Name)
		}

		// 如果字段是结构体，递归调用 structToMap
		if fieldValue.Kind() == reflect.Struct {
			result[fieldName] = StructToMap(fieldValue.Interface(), isSnake)
		} else if fieldValue.Kind() == reflect.Ptr && !fieldValue.IsNil() {
			// 如果字段是指针且不为 nil，递归调用 structToMap
			result[fieldName] = StructToMap(fieldValue.Elem().Interface(), isSnake)
		} else {
			// 否则，直接放入 map 中
			result[fieldName] = fieldValue.Interface()
		}
	}

	return result
}
