package common_util

import (
	"context"
	"fmt"
	"github.com/577683719/common/common/ctl"
	"github.com/577683719/common/common/util/common_util/date_util"
	common_util_snowflake "github.com/577683719/common/common/util/common_util/snowflake"
	"github.com/577683719/common/common/util/common_util/xid_util"
	"github.com/577683719/common/config"
	"github.com/duke-git/lancet/v2/netutil"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/google/uuid"
	"math/rand"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const allCharNum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const DATA = "2006-01-02"
const DATA_TIME = "2006-01-02 15:04:05"

func GenerateUniqueID(label string, num int) string {
	id := uuid.New()
	uniqueID := id.String()[:num]
	uniqueIDWithSID := label + uniqueID

	return uniqueIDWithSID
}

// 检验用户权限
func VerifyUserPermissions(ctx context.Context, userId int64) {
	info, err := ctl.GetUserInfo(ctx)
	if err != nil {
		panic("token错误")
	}
	if info.Id != userId {
		panic("用户没有操作权限")
	}
}
func VerifyUserPermissionRtuenUserId(ctx context.Context) int64 {
	info, err := ctl.GetUserInfo(ctx)
	if err != nil {
		panic("token错误")
	}
	return info.Id
}
func StructToMapStr(obj interface{}) map[string]string {
	objValue := reflect.ValueOf(obj)
	objType := objValue.Type()

	if objType.Kind() != reflect.Struct {
		panic("Invalid input: not a struct")
	}

	result := make(map[string]string)

	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		fieldValue := objValue.Field(i).Interface()
		tag := field.Tag.Get("json")

		// 如果没有指定json标签，则使用字段名称作为key
		if tag == "" {
			tag = field.Name
		}

		result[tag] = fmt.Sprint(fieldValue)
	}

	return result
}
func GetRandomCode(length int) string {
	sources := "0123456789"
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = sources[rand.Intn(len(sources))]
	}
	return string(result)
}

func GetSnowflakeId() string {
	return common_util_snowflake.GenIDStr()
}
func GetTransactionNumber33() string {
	var outTradeNo = fmt.Sprintf("%s%d", date_util.DateTimeToDateTimeNumberStr(time.Now()), xid_util.Next())
	return outTradeNo
}

func GetTransactionNumber29() string {
	return fmt.Sprintf("%s%s", date_util.DateTimeToDateTimeNumberStr(time.Now()), common_util_snowflake.GenIDStr())
}

// cps_bank_card   // 户名 银行账户  开户行   //
func GetTransactionNumber27() string {
	var outTradeNo = fmt.Sprintf("%s%d", date_util.DateTimeToDateNumberStr(time.Now()), xid_util.Next())
	return outTradeNo
}
func GetStringNumRandom(length int) string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	result := make([]byte, length)
	for i := range result {
		result[i] = allCharNum[r.Intn(len(allCharNum))]
	}
	return string(result)
}
func GetServiceAddr() string {
	ip := config.Conf.Server.Ip
	if ip == "" {
		ip = GetIP4()
	}
	return ip + ":" + config.Conf.Server.Port
}
func GetIP4() string {

	if runtime.GOOS == "windows" {
		ips := netutil.GetIps()
		for _, ip := range ips {
			prefixAny := strutil.HasPrefixAny(ip, []string{"192.168"})
			if prefixAny {
				return ip
			}
		}
	}
	ip := netutil.GetInternalIp()
	return ip
}

func IsValidDateFormat(input string) bool {
	_, err := time.Parse(DATA, input)
	return err == nil
}
func IsValidDateTimeFormat(input string) bool {
	_, err := time.Parse(DATA_TIME, input)
	return err == nil
}
func RandomNumber(min int, max int) int {
	// 使用当前时间的Unix时间戳作为随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生成 10 到 200 范围内的随机数
	randomNumber := rand.Intn(max-min+1) + min
	return randomNumber
}
func RandomNumberStr(min int, max int) string {
	// 使用当前时间的Unix时间戳作为随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生成 10 到 200 范围内的随机数
	randomNumber := rand.Intn(max-min+1) + min
	return strconv.Itoa(randomNumber)
}

func parseMinutesAndSeconds(durationStr string) (int, int, error) {
	parts := strings.Split(durationStr, "m")
	minutes, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}

	var secondsStr string
	// 检查's'前面是否有'.'
	if strings.Contains(parts[1], ".") {
		subParts := strings.Split(parts[1], ".")
		secondsStr = subParts[0]
	} else {
		subParts := strings.Split(parts[1], "s")
		if len(subParts) >= 1 {
			secondsStr = subParts[0]
		} else {
			return 0, 0, fmt.Errorf("无法解析秒数")
		}
	}

	seconds, err := strconv.Atoi(secondsStr)
	if err != nil {
		return 0, 0, err
	}

	return minutes, seconds, nil
}
func parseSeconds(durationStr string) (int, error) {
	parts := strings.Split(durationStr, "s")
	var secondsStr string
	// 检查's'前面是否有'.'
	if strings.Contains(parts[1], ".") {
		subParts := strings.Split(parts[1], ".")
		secondsStr = subParts[0]
	} else {
		subParts := strings.Split(parts[1], "s")
		if len(subParts) >= 1 {
			secondsStr = subParts[0]
		} else {
			return 0, fmt.Errorf("无法解析秒数")
		}
	}

	seconds, err := strconv.Atoi(secondsStr)
	if err != nil {
		return 0, err
	}

	return seconds, nil
}

// RemoveDecimalAndAddS 去掉小数点后面的值并加上s（如果有小数点）
func RemoveDecimalAndAddS(input string) string {
	// 查找小数点的位置
	dotIndex := strings.Index(input, ".")

	// 如果找到了小数点，则截取小数点前的部分并加上s
	if dotIndex != -1 {
		return input[:dotIndex] + "s"
	}

	// 如果没有小数点，保持原样
	return input
}

// 转成时分秒
func DurationToDaysHoursMinutesSeconds(durationStr string) (string, error) {
	durationStr = RemoveDecimalAndAddS(durationStr)
	parts := strings.Split(durationStr, "h")
	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		// 如果没有 "h"，则直接解析分钟和秒
		minutes, seconds, err := parseMinutesAndSeconds(durationStr)
		if err != nil {
			// 如果没有 "m"，则直接解析秒
			seconds, err := parseSeconds(durationStr)
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("%d秒", seconds), nil
		}
		return fmt.Sprintf("%d分%d秒", minutes, seconds), nil
	}

	// 计算天数
	days := hours / 24
	hours %= 24

	// 分钟和秒部分不需要处理，直接从原字符串中获取
	parts = strings.Split(parts[1], "m")
	minutes, err := strconv.Atoi(parts[0])
	if err != nil {
		return "", err
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
		result += fmt.Sprintf("%d小时", hours)
	}
	if minutes > 0 {
		result += fmt.Sprintf("%d分", minutes)
	}
	if seconds > 0 {
		result += fmt.Sprintf("%d秒", seconds)
	}

	return result, nil
}
func FailOnErr(format string, err error) {
	if err != nil {
		fmt.Printf(format, err)
	}
}

func Validata(flag bool, errStr string, a ...any) {
	if flag {
		fmt.Errorf(errStr, a)
	}
}
func GroupList[T any](data []T, chunkSize int) [][]T {
	if chunkSize <= 0 {
		return nil
	}

	var chunks [][]T
	dataLen := len(data)

	for i := 0; i < dataLen; i += chunkSize {
		end := i + chunkSize
		if end > dataLen {
			end = dataLen
		}
		chunks = append(chunks, data[i:end])
	}
	return chunks
}
func IntPtr(v int) *int {
	return &v
}

func Int64Ptr(v int64) *int64 {
	return &v
}

func UintPtr(v uint) *uint {
	return &v
}

func Uint64Ptr(v uint64) *uint64 {
	return &v
}

func Float64Ptr(v float64) *float64 {
	return &v
}

func BoolPtr(v bool) *bool {
	return &v
}

func StringPtr(v string) *string {
	return &v
}

// Helper function to convert a slice to a map for faster lookups
func sliceToMap(slice []string) map[string]struct{} {
	m := make(map[string]struct{})
	for _, v := range slice {
		m[v] = struct{}{}
	}
	return m
}

// Find differences between two string slices
func FindDifferences(a, b []string) (onlyInA, onlyInB []string) {
	mapA := sliceToMap(a)
	mapB := sliceToMap(b)

	for item := range mapA {
		if _, exists := mapB[item]; !exists {
			onlyInA = append(onlyInA, item)
		}
	}

	for item := range mapB {
		if _, exists := mapA[item]; !exists {
			onlyInB = append(onlyInB, item)
		}
	}

	return onlyInA, onlyInB
}
