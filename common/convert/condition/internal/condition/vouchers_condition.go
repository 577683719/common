package vouchers_condition

import (
	"github.com/577683719/common/service_api/repository/dao"
	"github.com/577683719/common/service_api/repository/model"
	"time"
)

func GetCondition(source model.Vouchers) dao.Condition {
	condition := dao.Condition{}
	condition.AndWithCondition(!determineIfItIsAValidValue(source.Category), model.VouchersColumns.Category, "=", source.Category)
	condition.AndWithCondition(!determineIfItIsAValidValue(source.Publish), model.VouchersColumns.Publish, "=", source.Publish)
	condition.AndWithCondition(!determineIfItIsAValidValue(source.VouchersImg), model.VouchersColumns.VouchersImg, "=", source.VouchersImg)
	condition.AndWithCondition(!determineIfItIsAValidValue(source.VouchersTitle), model.VouchersColumns.VouchersTitle, "=", source.VouchersTitle)
	condition.AndWithCondition(!determineIfItIsAValidValue(source.Price), model.VouchersColumns.Price, "=", source.Price)
	condition.AndWithCondition(!determineIfItIsAValidValue(source.UserLimit), model.VouchersColumns.UserLimit, "=", source.UserLimit)
	condition.AndWithCondition(!determineIfItIsAValidValue(source.StartTime), model.VouchersColumns.StartTime, "=", source.StartTime)
	condition.AndWithCondition(!determineIfItIsAValidValue(source.EndTime), model.VouchersColumns.EndTime, "=", source.EndTime)
	condition.AndWithCondition(!determineIfItIsAValidValue(source.PublishCount), model.VouchersColumns.PublishCount, "=", source.PublishCount)
	condition.AndWithCondition(!determineIfItIsAValidValue(source.Stock), model.VouchersColumns.Stock, "=", source.Stock)
	condition.AndWithCondition(!determineIfItIsAValidValue(source.CreateTime), model.VouchersColumns.CreateTime, "=", source.CreateTime)

	return condition
}

func GetInstanceCondition(source *model.Vouchers) dao.Condition {
	condition := dao.Condition{}
	condition.AndWithCondition(!determineIfItIsAValidValue(source.Category), model.VouchersColumns.Category, "=", source.Category)
	condition.AndWithCondition(!determineIfItIsAValidValue(source.Publish), model.VouchersColumns.Publish, "=", source.Publish)
	condition.AndWithCondition(!determineIfItIsAValidValue(source.VouchersImg), model.VouchersColumns.VouchersImg, "=", source.VouchersImg)
	condition.AndWithCondition(!determineIfItIsAValidValue(source.VouchersTitle), model.VouchersColumns.VouchersTitle, "=", source.VouchersTitle)
	condition.AndWithCondition(!determineIfItIsAValidValue(source.Price), model.VouchersColumns.Price, "=", source.Price)
	condition.AndWithCondition(!determineIfItIsAValidValue(source.UserLimit), model.VouchersColumns.UserLimit, "=", source.UserLimit)
	condition.AndWithCondition(!determineIfItIsAValidValue(source.StartTime), model.VouchersColumns.StartTime, "=", source.StartTime)
	condition.AndWithCondition(!determineIfItIsAValidValue(source.EndTime), model.VouchersColumns.EndTime, "=", source.EndTime)
	condition.AndWithCondition(!determineIfItIsAValidValue(source.PublishCount), model.VouchersColumns.PublishCount, "=", source.PublishCount)
	condition.AndWithCondition(!determineIfItIsAValidValue(source.Stock), model.VouchersColumns.Stock, "=", source.Stock)
	condition.AndWithCondition(!determineIfItIsAValidValue(source.CreateTime), model.VouchersColumns.CreateTime, "=", source.CreateTime)

	return condition
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
