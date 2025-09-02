package utils

import (
	"fmt"
	"time"
)

// CheckTimeFormat 检查时间格式是否正确
//
// @Description: 检查时间格式是否正确 1-年月日时分秒毫秒 2-年月日时分秒 3-年月日时分 4-年月日时 5-年月日 6-年月 7-年 8-时分秒 9-时分 10-时
// @param timeStr 时间字符串
// @param timeType 时间类型
// @return bool 正确返回true，错误返回false
func CheckTimeFormat(timeStr string, timeType int) bool {
	timeFormatTpl := "2006-01-02 15:04:05"
	switch timeType {
	case 1:
		timeFormatTpl = "2006-01-02 15:04:05.000"
		break
	case 2:
		timeFormatTpl = "2006-01-02 15:04:05"
		break
	case 3:
		timeFormatTpl = "2006-01-02 15:04"
		break
	case 4:
		timeFormatTpl = "2006-01-02 15"
		break
	case 5:
		timeFormatTpl = "2006-01-02"
		break
	case 6:
		timeFormatTpl = "2006-01"
		break
	case 7:
		timeFormatTpl = "2006"
		break
	case 8:
		timeFormatTpl = "15:04:05"
		break
	case 9:
		timeFormatTpl = "15:04"
		break
	case 10:
		timeFormatTpl = "15"
		break
	}
	_, err := time.Parse(timeFormatTpl, timeStr)
	if err != nil {
		return false
	}
	return true
}

// TimeFormat 时间格式化
//
// @Description: 时间格式化 1-年月日时分秒毫秒 2-年月日时分秒 3-年月日时分 4-年月日时 5-年月日 6-年月 7-年 8-时分秒 9-时分 10-时
// @param t 时间戳
// @param timeType 时间类型
// @return string 格式化后的时间字符串
func TimeFormat(t int64, timeType int) string {
	if t <= 0 {
		return ""
	}
	if len(fmt.Sprintf("%d", t)) == 10 {
		t *= 1000
	} else if len(fmt.Sprintf("%d", t)) == 13 {
	} else {
		return ""
	}
	currTime := time.UnixMilli(t)
	timeFormatTpl := "2006-01-02 15:04:05"
	switch timeType {
	case 1:
		timeFormatTpl = "2006-01-02 15:04:05.000"
		break
	case 2:
		timeFormatTpl = "2006-01-02 15:04:05"
		break
	case 3:
		timeFormatTpl = "2006-01-02 15:04"
		break
	case 4:
		timeFormatTpl = "2006-01-02 15"
		break
	case 5:
		timeFormatTpl = "2006-01-02"
		break
	case 6:
		timeFormatTpl = "2006-01"
		break
	case 7:
		timeFormatTpl = "2006"
		break
	case 8:
		timeFormatTpl = "15:04:05"
		break
	case 9:
		timeFormatTpl = "15:04"
		break
	case 10:
		timeFormatTpl = "15"
		break
	}
	return currTime.Format(timeFormatTpl)
}

// TimeParse 时间解析
//
// @Description: 时间解析 1-年月日时分秒毫秒 2-年月日时分秒 3-年月日时分 4-年月日时 5-年月日 6-年月 7-年 8-时分秒 9-时分 10-时
// @param timeStr 时间字符串
// @param timeType 时间类型
// @return int64 时间戳(毫秒)
func TimeParse(timeStr string, timeType int) int64 {
	timeFormatTpl := "2006-01-02 15:04:05"
	switch timeType {
	case 1:
		timeFormatTpl = "2006-01-02 15:04:05.000"
		break
	case 2:
		timeFormatTpl = "2006-01-02 15:04:05"
		break
	case 3:
		timeFormatTpl = "2006-01-02 15:04"
		break
	case 4:
		timeFormatTpl = "2006-01-02 15"
		break
	case 5:
		timeFormatTpl = "2006-01-02"
		break
	case 6:
		timeFormatTpl = "2006-01"
		break
	case 7:
		timeFormatTpl = "2006"
		break
	case 8:
		timeFormatTpl = "15:04:05"
		break
	case 9:
		timeFormatTpl = "15:04"
		break
	case 10:
		timeFormatTpl = "15"
		break
	}
	t, err := time.ParseInLocation(timeFormatTpl, timeStr, time.Local)
	if err != nil {
		return 0
	}
	return t.UnixMilli()
}

// GetYesterdayStartTime 获取昨天开始时间
//
// @Description: 获取昨天开始时间
// @return int64 时间戳(毫秒)
func GetYesterdayStartTime() int64 {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	return time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 0, 0, 0, 0, time.Local).Unix()
}

// GetYesterdayEndTime 获取昨天结束时间
//
// @Description: 获取昨天结束时间
// @return int64 时间戳(毫秒)
func GetYesterdayEndTime() int64 {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	return time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 23, 59, 59, 0, time.Local).Unix()
}
