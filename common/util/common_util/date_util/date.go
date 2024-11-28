package date_util

import (
	"fmt"
	"github.com/duke-git/lancet/v2/datetime"
	"time"
)

const (
	//"yyyy-mm-dd hh:mm:ss"
	YMDHMS = "yyyy-mm-dd hh:mm:ss"
	//yyyy-mm-dd
	YMD   = "yyyy-mm-dd"
	YMDHM = "yyyy-mm-dd hh:mm"

	YMDHMS_NUM = "20060102150405"
	YMD_NUM    = "20060102"
)

func init() {
	// 加载中国时区
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("错误：", err)
		return
	}
	// 设置全局时区
	fmt.Println("加载中国时区初始化成功")

	time.Local = location
}

// IsValidTimeFormat 校验字符串是否符合指定的时间格式
func IsValidTimeFormat(timeStr, layout string) bool {
	_, err := time.Parse(layout, timeStr)
	return err == nil
}

// 获取date字符串可以加减时间
func GetDateStr(day int64) string {
	return DateToDateStr(datetime.AddDay(time.Now(), day))
}

// 获取time字符串可以加减时间
func GetDateTimeStr(day int64) string {
	return DateToDateTimeStr(datetime.AddDay(time.Now(), day))
}

// 获取开始time字符串
func GetStartTimeStr(day int64) string {
	return datetime.AddDay(time.Now(), day).Format("2006-01-02") + " 00:00:00"
}

// 获取开始time字符串
func GetStartTime(day int64) time.Time {
	return datetime.BeginOfDay(datetime.AddDay(time.Now(), day))
}

// 获取结束time
func GetEndTime(day int64) time.Time {
	return datetime.EndOfDay(datetime.AddDay(time.Now(), day))
}

// 获取结束time字符串
func GetEndTimeStr(day int64) string {
	return datetime.AddDay(time.Now(), day).Format("2006-01-02") + " 23:59:59"
}

// 转成时间字符串
func GetDateFormat(date time.Time, format string) time.Time {
	toTime, _ := datetime.FormatStrToTime(datetime.FormatTimeToStr(date, format), format, "Asia/Shanghai")
	return toTime
}
func DateTimeToDateTimeNumberStr(date time.Time) string {
	formattedTime := date.Format(YMDHMS_NUM)
	return formattedTime
}
func DateTimeToDateNumberStr(date time.Time) string {
	formattedTime := date.Format(YMD_NUM)
	return formattedTime
}

// 转成时间字符串
func DateToDateTimeStr(date time.Time) string {

	return datetime.FormatTimeToStr(date, YMDHMS)
}

// 转成时间字符窜YMD格式
func DateToDateStr(date time.Time) string {
	return datetime.FormatTimeToStr(date, YMD)
}
func DateTimeStrToTime(str string) time.Time {
	toTime, _ := datetime.FormatStrToTime(str, YMDHMS, "Asia/Shanghai")
	return toTime
}
func IsValidDateFormat(input string) bool {
	_, err := time.Parse(YMD, input)
	return err == nil
}
func IsValidDateTimeFormat(input string) bool {
	_, err := time.Parse(YMDHMS, input)
	return err == nil
}
