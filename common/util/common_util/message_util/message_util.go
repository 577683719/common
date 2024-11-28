package message_util

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

type SmsModel struct {
	SecretId    string //SecretId
	SecretKey   string //SecretKey
	Host        string //主机
	Region      string //地区
	SmsSdkAppId string //短信appId
	SiganName   string //签名
	TemplateId  string //模板id
}

func SendSms(phone string, code []string, snsModel SmsModel) error {

	credential := common.NewCredential(
		snsModel.SecretId,
		snsModel.SecretKey,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = snsModel.Host
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := sms.NewClient(credential, snsModel.Region, cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := sms.NewSendSmsRequest()

	request.PhoneNumberSet = common.StringPtrs([]string{phone})
	request.SmsSdkAppId = common.StringPtr(snsModel.SmsSdkAppId)
	request.TemplateId = common.StringPtr(snsModel.TemplateId)
	request.SignName = common.StringPtr(snsModel.SiganName)
	request.TemplateParamSet = common.StringPtrs(code)

	// 返回的resp是一个SendSmsResponse的实例，与请求对象对应
	response, err := client.SendSms(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return nil
	}
	if err != nil {
		return err
	}
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())

	if *response.Response.SendStatusSet[0].Code == *common.StringPtr("Ok") {
		return nil
	}
	return errors.NewTencentCloudSDKError("50005", "短信发送失败", *response.Response.RequestId)
}
