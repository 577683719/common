package test

import (
	convert_gen_omit "ecs_cloud/common/convert/omit"
	"ecs_cloud/idl/gen/v1/coupons"
	"ecs_cloud/service_api/repository/model"
	"testing"
)

func TestGenVoucher(t *testing.T) {
	// 获取结构体的类型
	sourcesArrays := []any{&model.Vouchers{}}
	targetArrrays := []any{&coupons.VouchersResponse{}}
	methodData := convert_gen_omit.BuildMethodData(sourcesArrays, targetArrrays)
	templateData := convert_gen_omit.BuildTemplateData("user_omit", sourcesArrays, targetArrrays)
	templateData.SetPathName("../internal/omit/user_omit.go")
	templateData.SetMethodData(methodData)
	templateData.GenConevert()
}
