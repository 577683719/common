package Interceptor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/577683719/common/common/ctl"
	"github.com/577683719/common/common/e"
	cache_key "github.com/577683719/common/common/util/constant/cache"
	auth_filter "github.com/577683719/common/common/util/constant/filter"
	rpc_constant "github.com/577683719/common/common/util/constant/rpc"
	"github.com/577683719/common/common/util/gmd5"
	"github.com/577683719/common/common/util/gredis"
	"github.com/577683719/common/common/util/jwt"
	"github.com/577683719/common/common/util/logger"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/strutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"runtime"
	"strings"
)

type Validator interface {
	Validate() error
}

func ParameterValidation(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	// 把客户端传过来的结构体类型 退化成 实现了校验方法的类型
	if r, ok := req.(Validator); ok {
		if err = r.Validate(); err != nil {
			fmt.Println("服务端拦截器中的验证器校验错误-->" + err.Error())
			return nil, status.Error(e.ErrorParameterErrors, e.GetMsg(e.ErrorParameterErrors)+err.Error())
		}
	}
	resp, err = handler(ctx, req)
	if err != nil {
		logger.LogrusObj.Error("响应的数据异常:%s", err.Error())
	}

	logger.LogrusObj.Infof("响应的数据是:%s", strutil.Substring(convertor.ToString(resp), 0, 1024))
	return resp, err
}
func logPanicInfo(info interface{}) (string, int) {
	// 获取堆栈信息

	// 获取调用函数信息
	pc, _, line, ok := runtime.Caller(5)
	var funcName string
	if ok {
		fn := runtime.FuncForPC(pc)
		funcName = fn.Name()
	}

	return funcName, line
	// 打印日志
}

func UnaryInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	method := info.FullMethod
	logger.LogrusObj.Infof("请求的路径是%s", method)
	bytesData, _ := json.Marshal(req)
	logger.LogrusObj.Infof("请求的参数是:%s", string(bytesData))
	// 创建一个新的上下文，用于传递错误对象
	ctx = context.WithValue(ctx, "grpcError", nil)
	// 使用 defer + recover() 捕获异常
	defer func() {
		if r := recover(); r != nil {
			// 获取堆栈信息
			stackBuf := make([]byte, 1024)
			stackBuf = stackBuf[:runtime.Stack(stackBuf, false)]
			// 这里你可以选择如何处理 err，例如打印到日志或者返回给客户端
			fmt.Println(status.Errorf(codes.Internal, "服务器异常: %v\n堆栈信息:\n%s", r, stackBuf))
			// 返回一个 gRPC 错误响应
			err = status.Errorf(e.ERROR, "服务器异常: %v", r)
			// 直接返回错误响应给客户端
		}
	}()

	//放行过滤的请求
	if auth_filter.UserFilterPaths[method] {
		return handler(ctx, req)
	}

	//进行权限认证
	ctx, err = auth(ctx)

	if err != nil {
		return checkRpcOrAdmin(ctx, req, handler, err)
	}

	return handler(ctx, req)
}

// 校验是否是rpc或者是admin
func checkRpcOrAdmin(ctx context.Context, req any, handler grpc.UnaryHandler, err error) (any, error) {
	//rpcoken
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, status.Errorf(400, "无法获取元数据")
	}

	if len(md["rpc-token"]) > 0 {
		if md["rpc-token"][0] == rpc_constant.RPC_TOKEN {
			//rpc请求进行放行不需要
			return handler(ctx, req)
		}
		//rpc请求进行放行不需要
		return nil, status.Error(500, "请求出错")
	}

	if len(md["admin-service"]) > 0 {
		if md["admin-service"][0] == rpc_constant.ADMIN_SERVICE_TOKEN {
			//rpc请求进行放行不需要
			return handler(ctx, req)
		}
		//rpc请求进行放行不需要
		return nil, status.Error(500, "请求出错")
	}
	return nil, err
}

func StreamInterceptor(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {

	method := info.FullMethod
	//放行过滤的请求
	if auth_filter.UserFilterPaths[method] {
		return handler(srv, ss)
	}
	//用户权限验证
	_, err := auth(ss.Context())

	if err != nil {
		return err
	}
	err = handler(srv, ss)

	return handler(srv, ss)
}

// 权限认证
func auth(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, status.Errorf(400, "无法获取元数据")
	}
	authorization := md["authorization"]
	if len(authorization) < 1 {
		return ctx, status.Errorf(e.ErrorAuth, e.GetMsg(e.ErrorAuth))
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")

	if token == "" {
		return ctx, status.Errorf(e.ErrorAuth, e.GetMsg(e.ErrorAuth))
	}
	claims, err := jwt.ParseToken(token)
	if err != nil {
		return ctx, status.Errorf(e.ErrorAuth, e.GetMsg(e.ErrorAuth))
	}
	//退出登录的验证
	//err = logoutVerification(claims, token)
	//if err != nil {
	//	return ctx, err
	//}
	// 验证成功后储存

	ctx = context.WithValue(ctx, "userInfo", &ctl.UserInfo{
		Id:        claims.UserID,
		AccountNo: claims.AccountNo,
		Username:  claims.UserName,
		Phone:     claims.Phone,
		Email:     claims.Email,
		Token:     token,
		Nickname:  claims.Nickname,
		ExpiresAt: claims.ExpiresAt,
	})

	return ctx, nil
}

// 退出登录的验证
func logoutVerification(claims *jwt.Claims, token string) error {
	logoutUser, _ := gredis.GetString(fmt.Sprintf(cache_key.UserLogout, claims.UserID, gmd5.Md5(token)))

	if logoutUser != "" {
		//账号已经退出登录
		return status.Error(e.ErrorAuthToken, e.GetMsg(e.ErrorAuthToken))
	}
	return nil
}
func requestsForRelease() {

}
