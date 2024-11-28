package cache_key

const (
	//用户退出登录
	UserLogout = "user:louout:%v:%s"
	//用户手机号验证吗
	SmsCode = "user:code:%s"

	//用户图形验证码y轴的值
	CaptchaCode = "user:captcha:code:%s"

	//用户邮箱验证码
	EmailCode = "user:code:%s"
	//优惠卷锁
	LockCoupon = "lock:coupon:%s"
	//用户关机锁
	LockStopDevice  = "lock:stop-device:%s"
	LockStartDevice = "lock:start-device:%s"

	ProductLeftNum = "product:leftNum:%s"

	//预支付key
	PayPreCreate           = "pay-service:pre-create:"
	IdentityAuthentication = "user:authentication:%d"
)
