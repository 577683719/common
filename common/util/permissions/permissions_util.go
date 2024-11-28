package permissions_util

import (
	"context"
	rpc_constant "ecs_cloud/common/util/constant/rpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// 校验是否可以访问admin接口的权限
func VerifyAdminPermissions(ctx context.Context) error {
	//rpcoken
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(400, "无法获取元数据")
	}
	if len(md["admin-service"]) < 1 || len(md["username"]) < 1 || len(md["userid"]) < 1 {
		return status.Error(500, "没有权限")
	}
	if !(md["admin-service"][0] == rpc_constant.ADMIN_SERVICE_TOKEN) {
		//rpc请求进行放行不需要
		return status.Error(500, "没有权限")
	}
	return nil
}
func GetAdminUserId(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(400, "无法获取元数据")
	}
	if len(md["admin-service"]) < 1 || len(md["username"]) < 1 || len(md["userid"]) < 1 {
		return "", status.Error(500, "没有权限")
	}
	if !(md["admin-service"][0] == rpc_constant.ADMIN_SERVICE_TOKEN) {
		//rpc请求进行放行不需要
		return "", status.Error(500, "没有权限")
	}
	return md["userid"][0], nil
}
func GetLable(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(400, "无法获取元数据")
	}
	if len(md["admin-service"]) < 1 || len(md["username"]) < 1 || len(md["userid"]) < 1 {
		return "", status.Error(500, "没有权限")
	}
	if !(md["admin-service"][0] == rpc_constant.ADMIN_SERVICE_TOKEN) {
		//rpc请求进行放行不需要
		return "", status.Error(500, "没有权限")
	}
	return md["supplier"][0], nil
}
