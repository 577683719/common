package main

import (
	"ecs_cloud/common/util/common_util"
	"fmt"
	"github.com/duke-git/lancet/v2/datetime"
	"strconv"
	"time"

	"strings"
)

func main() {
	//original := stream.FromSlice([]int{1, 2, 3})
	//
	//addOne := func(n int) int {
	//	return n + 1
	//}
	//
	//increament := original.Map(addOne)
	//
	//fmt.Println(increament.ToSlice())
	difference := datetime.AddDay(time.Now(), 30).Sub(time.Now())
	result, _ := common_util.DurationToDaysHoursMinutesSeconds(difference.String())
	fmt.Println(result)
	// Output:
	// [2 3 4]
}
func durationToDaysHoursMinutesSeconds(durationStr string) (string, error) {
	parts := strings.Split(durationStr, "h")
	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		// 如果没有 "h"，则直接解析分钟和秒
		minutes, seconds, err := parseMinutesAndSeconds(durationStr)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%dm%ds", minutes, seconds), nil
	}

	// 计算天数
	days := hours / 24
	hours %= 24

	// 分钟和秒部分不需要处理，直接从原字符串中获取
	parts = strings.Split(parts[1], "m")
	minutes, err := strconv.Atoi(parts[0])
	if err != nil {
		// 如果没有 "m"，则直接解析秒
		seconds, err := strconv.Atoi(strings.TrimSuffix(parts[0], "s"))
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%ds", seconds), nil
	}

	parts = strings.Split(parts[1], "s")
	seconds, err := strconv.Atoi(parts[0])
	if err != nil {
		return "", err
	}

	// 构建返回字符串
	var result string
	if days > 0 {
		result += fmt.Sprintf("%d天", days)
	}
	if hours > 0 {
		result += fmt.Sprintf("%d时", hours)
	}
	if minutes > 0 {
		result += fmt.Sprintf("%d分", minutes)
	}
	if seconds > 0 {
		result += fmt.Sprintf("%d秒", seconds)
	}

	return result, nil
}

func parseMinutesAndSeconds(durationStr string) (int, int, error) {
	parts := strings.Split(durationStr, "m")
	minutes, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}

	parts = strings.Split(parts[1], "s")
	seconds, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}

	return minutes, seconds, nil
}

func main1() {
	durationStr := "55h0m0s"
	result, err := common_util.DurationToDaysHoursMinutesSeconds(durationStr)
	if err != nil {
		fmt.Println("发生错误:", err)
	} else {
		fmt.Printf("%s 转换结果为: %s\n", durationStr, result)
	}
}
