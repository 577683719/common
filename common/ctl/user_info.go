package ctl

import (
	"context"
	"errors"
)

type key int

var userKey key

type UserInfo struct {
	Id        int64
	AccountNo string
	Username  string
	HeadImg   string
	Email     string
	Phone     string
	Token     string
	ExpiresAt int64
	Nickname  string
}

func GetUserName(ctx context.Context) (string, error) {
	userInfo, err := GetUserInfo(ctx)
	if err != nil {
		return "", err
	}
	return userInfo.Username, err
}
func GetUserId(ctx context.Context) (int64, error) {
	userInfo, err := GetUserInfo(ctx)
	if err != nil {
		return 0, err
	}
	return userInfo.Id, err
}
func GetUserInfo(ctx context.Context) (*UserInfo, error) {
	user, ok := FromContext(ctx)
	if !ok {
		return nil, errors.New("获取用户信息错误")
	}
	return user, nil
}

func NewContext(ctx context.Context, u *UserInfo) context.Context {
	return context.WithValue(ctx, userKey, u)
}

func FromContext(ctx context.Context) (*UserInfo, bool) {
	if ctx.Value("userInfo") == nil {
		return nil, false
	}
	info := ctx.Value("userInfo").(*UserInfo)
	return info, true
}

func InitUserInfo(ctx context.Context) {
	// TOOD 放缓存，之后的用户信息，走缓存
}
