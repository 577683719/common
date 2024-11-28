package etcd_util

import (
	"context"
	"encoding/json"
	"github.com/577683719/common/common/util/logger"
	"github.com/577683719/common/config"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/rand"
)

var (
	EtcdClient *clientv3.Client
)

type Server struct {
	Name    string `json:"name"`
	Addr    string `json:"addr"`    // 地址
	Version string `json:"version"` // 版本
	Weight  int64  `json:"weight"`  // 权重
}

func InitEtcdClient() *clientv3.Client {
	var err error
	EtcdClient, err = clientv3.New(clientv3.Config{
		Endpoints: []string{config.Conf.Etcd.Address}, // etcd 服务器地址
	})
	if err != nil {
		logger.LogrusObj.Errorf("etcd错误:%v", err)
	}
	return EtcdClient
}
func ParseValue(value []byte) (*Server, error) {
	server := &Server{}
	if err := json.Unmarshal(value, &server); err != nil {
		return server, err
	}

	return server, nil
}
func getServiceEndpoint(serviceName string) (*Server, error) {
	resp, err := EtcdClient.Get(context.Background(), serviceName, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	if len(resp.Kvs) == 0 {
		return nil, status.Error(codes.NotFound, "Service not found")
	}
	value := resp.Kvs[rand.Intn(len(resp.Kvs))].Value
	return ParseValue(value)
}
