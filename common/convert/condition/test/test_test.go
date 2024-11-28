package test

import (
	condition_gen "github.com/577683719/common/common/convert/condition"
	"github.com/577683719/common/idl/gen/v1/coupons"
	"github.com/577683719/common/service_api/repository/model"
	"testing"
)

func TestGenVoucher(t *testing.T) {
	// 获取结构体的类型
	sourcesArrays := []any{&model.Vouchers{}}
	targetArrrays := []any{&coupons.VouchersResponse{}}
	methodData := condition_gen.BuildMethodData(sourcesArrays, targetArrrays)
	templateData := condition_gen.BuildTemplateData("vouchers_condition", sourcesArrays, targetArrrays)
	templateData.SetPathName("../internal/condition/vouchers_condition.go")
	templateData.SetMethodData(methodData)
	templateData.GenConevert()
}
