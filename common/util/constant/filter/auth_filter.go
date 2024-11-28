package auth_filter

var (
	UserFilterPaths = map[string]bool{
		"/UserService/UserLogin":                  true,
		"/UserService/UserRegister":               true,
		"/UserService/EditPassword":               true,
		"/UserService/WechatLoginCallback":        true,
		"/UserService/GetWeChatLoginQRCode":       true,
		"/BillServices/ExportDailyStatement":      true,
		"/MessageService/SendSms":                 true,
		"/UserService/GetSlidingVerificationCode": true,
		"/UserService/WeChatBindingCallback":      true,
		"/CouponsService/HomePage":                true,
		"/UserService/BindPhone":                  true,
	}
)
