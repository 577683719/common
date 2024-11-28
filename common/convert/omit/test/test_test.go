package test

import (
	convert_gen_omit "github.com/577683719/common/common/convert/omit"
	"github.com/577683719/common/idl/gen/v1/coupons"
	"github.com/577683719/common/service_api/repository/model"
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
