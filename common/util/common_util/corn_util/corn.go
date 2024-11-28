package corn_util

import (
	"fmt"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/datetime"
	"strings"
	"time"
)

// GenerateCronExpression 生成一个 Cron 表达式，表示在指定时间后执行一次
func DateTimeToCronStr(executionTime time.Time) string {
	// 将 JobFunc 包装成一个 Cron Job
	cron := fmt.Sprintf("%d %d %d %d %d ? %d-%d ", executionTime.Second(), executionTime.Minute(), executionTime.Hour(), executionTime.Day(), executionTime.Month(), executionTime.Year(), executionTime.Year())
	return cron
}
func GetCornDateTimeStr(cron string) string {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	cornStr := strings.Split(cron, " ")
	year, _ := convertor.ToInt(strings.TrimSpace(strings.Split(cron, "-")[1]))
	month, _ := convertor.ToInt(strings.TrimSpace(cornStr[4]))
	day, _ := convertor.ToInt(strings.TrimSpace(cornStr[3]))
	hour, _ := convertor.ToInt(strings.TrimSpace(cornStr[2]))
	minute, _ := convertor.ToInt(strings.TrimSpace(cornStr[1]))
	second, _ := convertor.ToInt(strings.TrimSpace(cornStr[0]))
	executionTime := time.Date(int(year), time.Month(month), int(day), int(hour), int(minute), int(second), 0, loc)
	date := datetime.AddDay(executionTime, 0)
	return datetime.FormatTimeToStr(date, "yyyy-mm-dd hh:mm:ss")
}

// 添加一天
func AddCornDay(cron string, dayCount int64) string {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	cornStr := strings.Split(cron, " ")
	year, _ := convertor.ToInt(strings.TrimSpace(strings.Split(cron, "-")[1]))
	month, _ := convertor.ToInt(strings.TrimSpace(cornStr[4]))
	day, _ := convertor.ToInt(strings.TrimSpace(cornStr[3]))
	hour, _ := convertor.ToInt(strings.TrimSpace(cornStr[2]))
	minute, _ := convertor.ToInt(strings.TrimSpace(cornStr[1]))
	second, _ := convertor.ToInt(strings.TrimSpace(cornStr[0]))
	executionTime := time.Date(int(year), time.Month(month), int(day), int(hour), int(minute), int(second), 0, loc)
	date := datetime.AddDay(executionTime, dayCount)
	return DateTimeToCronStr(date)
}

// 添加月
func AddCornMonth(cron string, minuteCount int64) string {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	cornStr := strings.Split(cron, " ")
	year, _ := convertor.ToInt(strings.TrimSpace(strings.Split(cron, "-")[1]))
	month, _ := convertor.ToInt(strings.TrimSpace(cornStr[4]))
	day, _ := convertor.ToInt(strings.TrimSpace(cornStr[3]))
	hour, _ := convertor.ToInt(strings.TrimSpace(cornStr[2]))
	minute, _ := convertor.ToInt(strings.TrimSpace(cornStr[1]))
	second, _ := convertor.ToInt(strings.TrimSpace(cornStr[0]))
	executionTime := time.Date(int(year), time.Month(month), int(day), int(hour), int(minute), int(second), 0, loc)
	date := datetime.AddDay(executionTime, minuteCount*30)
	return DateTimeToCronStr(date)
}

// 添加一分钟
func AddCornMinute(cron string, minuteCount int64) string {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	cornStr := strings.Split(cron, " ")
	year, _ := convertor.ToInt(strings.TrimSpace(strings.Split(cron, "-")[1]))
	month, _ := convertor.ToInt(strings.TrimSpace(cornStr[4]))
	day, _ := convertor.ToInt(strings.TrimSpace(cornStr[3]))
	hour, _ := convertor.ToInt(strings.TrimSpace(cornStr[2]))
	minute, _ := convertor.ToInt(strings.TrimSpace(cornStr[1]))
	second, _ := convertor.ToInt(strings.TrimSpace(cornStr[0]))
	executionTime := time.Date(int(year), time.Month(month), int(day), int(hour), int(minute), int(second), 0, loc)
	date := datetime.AddMinute(executionTime, minuteCount)
	return DateTimeToCronStr(date)
}

// 添加几周
func AddCornWeek(cron string, weekCount int64) string {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	cornStr := strings.Split(cron, " ")
	year, _ := convertor.ToInt(strings.TrimSpace(strings.Split(cron, "-")[1]))
	month, _ := convertor.ToInt(strings.TrimSpace(cornStr[4]))
	day, _ := convertor.ToInt(strings.TrimSpace(cornStr[3]))
	hour, _ := convertor.ToInt(strings.TrimSpace(cornStr[2]))
	minute, _ := convertor.ToInt(strings.TrimSpace(cornStr[1]))
	second, _ := convertor.ToInt(strings.TrimSpace(cornStr[0]))
	executionTime := time.Date(int(year), time.Month(month), int(day), int(hour), int(minute), int(second), 0, loc)
	date := datetime.AddDay(executionTime, weekCount*7)
	return DateTimeToCronStr(date)
}
