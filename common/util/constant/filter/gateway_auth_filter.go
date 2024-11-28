package auth_filter

var (
	AccountServiceFilterPaths = map[string]bool{
		"/api/v1/account/login":                      true,
		"/api/v1/account/register":                   true,
		"/api/v1/account/loginout":                   true,
		"/api/v1/account/editPassword":               true,
		"/api/v1/message/sendSms":                    true,
		"/api/pay/ali/notify":                        true,
		"/api/pay/wechat/notify":                     true,
		"/api/v1/account/getWeChatLoginQRCode":       true,
		"/api/v1/account/wechatLoginCallback":        true,
		"/api/v1/account/getSlidingVerificationCode": true,
		"/api/v1/account/weChatBindingCallback":      true,
		"/api/v1/activity/homePage":                  true,
		"/api/pay/ali/integralNotify":                true,
		"/api/pay/ali/walletNotify":                  true,
		"/api/pay/wechat/walletNotify":               true,
		"/api/pay/wechat/integralNotify":             true,
		"/api/v1/account/bindPhone":                  true,
	}
)
