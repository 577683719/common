package user_omit

import (
	"github.com/577683719/common/service_api/repository/model"
	"time"
)

func GetOmit(source model.Vouchers) []string {
	var str []string

	if determineIfItIsAValidValue(source.Category) {
		str = append(str, model.VouchersColumns.Category)
	}
	if determineIfItIsAValidValue(source.Publish) {
		str = append(str, model.VouchersColumns.Publish)
	}
	if determineIfItIsAValidValue(source.VouchersImg) {
		str = append(str, model.VouchersColumns.VouchersImg)
	}
	if determineIfItIsAValidValue(source.VouchersTitle) {
		str = append(str, model.VouchersColumns.VouchersTitle)
	}
	if determineIfItIsAValidValue(source.Price) {
		str = append(str, model.VouchersColumns.Price)
	}
	if determineIfItIsAValidValue(source.UserLimit) {
		str = append(str, model.VouchersColumns.UserLimit)
	}
	if determineIfItIsAValidValue(source.StartTime) {
		str = append(str, model.VouchersColumns.StartTime)
	}
	if determineIfItIsAValidValue(source.EndTime) {
		str = append(str, model.VouchersColumns.EndTime)
	}
	if determineIfItIsAValidValue(source.PublishCount) {
		str = append(str, model.VouchersColumns.PublishCount)
	}
	if determineIfItIsAValidValue(source.Stock) {
		str = append(str, model.VouchersColumns.Stock)
	}
	if determineIfItIsAValidValue(source.CreateTime) {
		str = append(str, model.VouchersColumns.CreateTime)
	}

	return str
}

// DetermineIfItIsAValidValue 判断不同类型的值是否有效
func determineIfItIsAValidValue(value interface{}) bool {
	// 根据需要进行类型转换的逻辑
	switch v := value.(type) {
	case time.Time:
		return v.IsZero() // 时间类型不为零则有效
	case string:
		return v == "" // 字符串非空则有效
	case int, int8, int16, int32, int64:
		return v == 0 // 整数不为零则有效
	case uint, uint8, uint16, uint32, uint64:
		return v == 0 // 无符号整数不为零则有效
	case float32, float64:
		return v == 0.0 // 浮点数不为零则有效
	default:
		return false // 其他类型默认不统计
	}
}
