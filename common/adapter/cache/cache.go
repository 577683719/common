package cache_adapter

import (
	"encoding/json"
	"fmt"
	"github.com/577683719/common/common/util/gredis"
	"github.com/577683719/common/idl/gen/v1/enums"
	"github.com/patrickmn/go-cache"
	"time"
)

func Set(cache, key string, value interface{}) {
	switch cache {
	case enums.CacheAdapterEnums_REDIS.String():
		adapter := RedisAdapter{}
		adapter.Set(key, value)
	}
}
func Get(cache, key string) string {
	switch cache {
	case enums.CacheAdapterEnums_REDIS.String():
		adapter := RedisAdapter{}
		return adapter.Get(key)
	}
	return ""
}
func SetExpiration(cache, key string, value interface{}, expiration time.Duration) {
	switch cache {
	case enums.CacheAdapterEnums_REDIS.String():
		redisAdapter := RedisAdapter{}
		redisAdapter.SetExpiration(key, value, expiration)
	}
}

type CacheAdapter struct {
	cache *cache.Cache
}

type RedisAdapter struct {
}

func (adapter *RedisAdapter) SetExpiration(key string, value interface{}, expiration time.Duration) string {
	marshal, _ := json.Marshal(value)
	err := gredis.SetExpiration(key, string(marshal), expiration)
	fmt.Println(err)
	return string(marshal)
}

func (adapter *RedisAdapter) Set(key string, value interface{}) string {
	marshal, _ := json.Marshal(value)
	gredis.SetString(key, string(marshal))
	return string(marshal)
}

func (adapter *RedisAdapter) Get(key string) string {
	value, _ := gredis.GetString(key)
	return value
}

type StdCache interface {
	Get(key string) string
	Set(key string, value interface{}) string
	SetExpiration(key string, value interface{}, expiration time.Duration) string
}
