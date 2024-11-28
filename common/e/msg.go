package e

var MsgFlags = map[int]string{
	SUCCESS: "ok",
	ERROR:   "fail",

	InvalidParams:              "请求参数错误",
	HaveSignUp:                 "已经报名了",
	ErrorActivityTimeout:       "活动过期了",
	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token已超时",
	ErrorAuthToken:             "Token失效",
	ErrorAuth:                  "Token错误",
	ErrorNotCompare:            "不匹配",
	ErrorDatabase:              "数据库操作出错,请重试",
	ErrorAuthNotFound:          "Token不能为空",

	ErrorServiceUnavailable: "过载保护，服务暂时不可用",
	ErrorDeadline:           "服务调用超时",
	ErrorSMSCode:            "短信码错误",
	ErrorAccountExistence:   "账号已存在",
	ErrorEmialExistence:     "邮箱已经存在",
	ErrorAccountErr:         "账号或密码错误",
	ErrorPhoneErr:           "手机号或密码错误",

	ErrorBusinessErrors:    "业务异常",
	ErrorCreateOrderErrors: "创建订单失败",
	ErrorOrderAuth:         "创建订单token校验失败",
	ErrorParameterErrors:   "请求参数错误",
	ErrorOrderList:         "获取订单列表失败",
	ErrorVouchersList:      "获取代金券列表失败",
	ErrorUsersList:         "获取用户列表失败",

	ErrorNoPermissions:              "用户没有操作权限",
	ErrorPickUp:                     "领取失败",
	ErrorGetUserInfo:                "获取用户信息失败",
	ErroBalanceIsInsufficientErrors: "用户余额不足",
	ErrorProductList:                "获取产品列表失败",

	ErrorStartInstanceErrors:   "开启实例失败",
	ErrorReleaseInstanceErrors: "釋放实例失败",
	ErrorStopInstanceErrors:    "关闭实例失败",
	ErrorAddVouchers:           "添加代金券失败",
	ErrorOrderProductLeftNum:   "库存扣减失败",
}

// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
