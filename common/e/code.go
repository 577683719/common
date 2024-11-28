package e

const (
	SUCCESS       = 200
	ERROR         = 500
	InvalidParams = 400

	// 成员错误
	ErrorExistUser      = 10002
	ErrorNotExistUser   = 10003
	ErrorFailEncryption = 10006
	ErrorNotCompare     = 10007

	HaveSignUp           = 20001
	ErrorActivityTimeout = 20002

	ErrorAuthCheckTokenFail    = 30001 // token 错误
	ErrorAuthCheckTokenTimeout = 30002 // token 过期
	ErrorAuthToken             = 30003
	ErrorAuth                  = 30004
	ErrorAuthNotFound          = 30005
	ErrorDatabase              = 40001

	ErrorServiceUnavailable = 50003
	ErrorDeadline           = 50004

	// 登录相关
	ErrorSMSCode          = 60001
	ErrorEmialCode        = 60002
	ErrorAccountExistence = 60003
	ErrorEmialExistence   = 60004
	ErrorAccountErr       = 60005
	ErrorPhoneErr         = 60006
	ErrorEmailErr         = 60007

	// 用户相关
	ErrorNoPermissions              = 70000
	ErrorPickUp                     = 70001
	ErrorGetUserInfo                = 70002
	ErroBalanceIsInsufficientErrors = 70003
	ErrorOrderList                  = 70004
	ErrorVouchersList               = 70005
	ErrorUsersList                  = 70006

	// 业务相关
	ErrorBusinessErrors    = 80000
	ErrorCreateOrderErrors = 800001
	ErrorOrderAuth         = 800002
	ErrorParameterErrors   = 800003
	ErrorProductList       = 80004

	ErrorStartInstanceErrors   = 90000
	ErrorStopInstanceErrors    = 90001
	ErrorReleaseInstanceErrors = 90002

	ErrorAddVouchers = 900010

	// 订单相关
	ErrorOrderProductLeftNum = 900100
)
