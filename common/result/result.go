package result

import (
	"ecs_cloud/common/e"
)

// Result 表示统一响应的JSON格式
type Result struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 响应消息
	Data    interface{} `json:"data"`    // 响应数据
}

// 异常信息从定义好的bizerr里面获取
func Fail(code int) *Result {
	message := e.GetMsg(code)
	return &Result{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}

// 程序异常
func Err(code int, message string) *Result {
	return &Result{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}

// 接口执行正常 需要返回数据 data
func Ok(code int, message string, data interface{}) *Result {
	return &Result{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// 请求成功返回data
func Data(data interface{}) *Result {
	return &Result{
		Code:    e.SUCCESS,
		Message: e.GetMsg(e.SUCCESS),
		Data:    data,
	}
}

// 请求成功返回封装好的bizcode
func Code(code int) *Result {
	message := e.GetMsg(code)
	return &Result{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}
