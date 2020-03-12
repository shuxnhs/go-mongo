package time

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// 获取unix时间戳
func GetTimestamp() int64 {
	return time.Now().Unix()
}

// 格式化日期 Y-m-d H:i:s
func Date(format string, timestamp int64) string {
	var t time.Time

	if timestamp > 0 {
		t = time.Unix(timestamp, 0)
	} else {
		t = time.Now()
	}

	// 替换年月日
	format = strings.Replace(format, "Y", fmt.Sprintf("%d", t.Year()), -1)
	format = strings.Replace(format, "m", fmt.Sprintf("%02d", int(t.Month())), -1)
	format = strings.Replace(format, "d", fmt.Sprintf("%02d", t.Day()), -1)

	// 替换时分秒
	format = strings.Replace(format, "H", fmt.Sprintf("%02d", t.Hour()), -1)
	format = strings.Replace(format, "i", fmt.Sprintf("%02d", t.Minute()), -1)
	format = strings.Replace(format, "s", fmt.Sprintf("%02d", t.Second()), -1)

	return format
}

//字符串转时间戳
func StrToTimestamp(str string) int64 {
	format := [4]string{
		"20060102",
		"2006-01-02",
		"2006-01-02 15:04",
		"2006-01-02 15:04:05",
	}

	for i := 0; i < 4; i++ {
		t, err := time.ParseInLocation(format[i], str, time.Local)
		if err == nil {
			timestamp := t.Unix()
			return timestamp
		}
	}
	return 0
}

func GetDate() string {
	return time.Now().Format("2006-01-02")
}

func GetDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetDay() int {
	day, _ := strconv.Atoi(time.Now().Format("02"))
	return day
}
