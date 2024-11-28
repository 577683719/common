package test

import (
	condition_gen "ecs_cloud/common/convert/condition"
	"ecs_cloud/idl/gen/v1/coupons"
	"ecs_cloud/service_api/repository/model"
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
