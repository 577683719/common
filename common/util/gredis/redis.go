package gredis

import (
	"context"
	"ecs_cloud/common/util/common_util/date_util"
	"ecs_cloud/config"
	"fmt"
	"github.com/duke-git/lancet/v2/datetime"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	Rdb     *redis.Client
	Ctx     = context.Background()
	Redsync *redsync.Redsync
)

func InitRedis() (err error) {
	redisConfig := config.Conf.Redis

	Rdb = redis.NewClient(&redis.Options{
		Addr:         redisConfig.Address,
		Password:     redisConfig.Password, // no password set
		DB:           redisConfig.Db,       // use default DB
		PoolSize:     32,                   //最大连接数
		MinIdleConns: 16,                   //最小连接锁
		PoolTimeout:  30 * time.Second,     //超时时间
	})

	_, err = Rdb.Ping(Ctx).Result()
	if err != nil {
		return err
	}
	//初始化分布式锁
	pool := goredis.NewPool(Rdb)
	Redsync = redsync.New(pool)
	fmt.Println("redis初始化成功时间:", datetime.FormatTimeToStr(time.Now(), date_util.YMDHMS))
	return nil
}
func Incr(key string) (int64, error) {
	return Rdb.Incr(Ctx, key).Result()
}
func IncrBy(key string, value int64) (int64, error) {
	return Rdb.IncrBy(Ctx, key, value).Result()
}
func Decr(key string) (int64, error) {
	return Rdb.Decr(Ctx, key).Result()
}

// 设置缓存
func Set(key string, value interface{}, expiration time.Duration) (err error) {
	err = Rdb.Set(Ctx, key, value, expiration).Err()
	return nil
}

// 设置缓存并且有过期时间
func SetExpiration(key string, value interface{}, expiration time.Duration) (err error) {
	err = Rdb.Set(Ctx, key, value, expiration).Err()
	return nil
}

// SetString 设置字符串类型的键值对，不过期
func SetString(key, value string) error {
	return Rdb.Set(Ctx, key, value, 0).Err()
}

// ExecLuaScript 执行 Lua 脚本操作 Redis
func ExecLuaScript(ctx context.Context, script string, keys []string, args ...interface{}) (interface{}, error) {
	result, err := Rdb.Eval(ctx, script, keys, args...).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SetStringWithExp 设置字符串类型的键值对，带过期时间（秒）
func SetStringWithExp(key, value string, expiration time.Duration) error {
	return Rdb.Set(Ctx, key, value, expiration).Err()
}

// GetString 获取字符串类型的键的值
func GetString(key string) (string, error) {
	return Rdb.Get(Ctx, key).Result()
}

// HSet 设置哈希字段的字符串值，不过期
func HSet(key, field, value string) error {
	return Rdb.HSet(Ctx, key, field, value).Err()
}

// HSetWithExp 设置哈希字段的字符串值，带过期时间（秒）
func HSetWithExp(key, field, value string, expiration time.Duration) error {
	return Rdb.HSet(Ctx, key, field, value).Err()
}

// HGet 获取哈希字段的字符串值
func HGet(key, field string) (string, error) {
	return Rdb.HGet(Ctx, key, field).Result()
}

// LPush 将一个或多个值插入到列表头部，不过期
func LPush(key string, values ...interface{}) (int64, error) {
	return Rdb.LPush(Ctx, key, values...).Result()
}

// LPushWithExp 将一个或多个值插入到列表头部，带过期时间（秒）
func LPushWithExp(key string, expiration time.Duration, values ...interface{}) (int64, error) {
	pipe := Rdb.Pipeline()
	cmd := pipe.LPush(Ctx, key, values...)
	pipe.Expire(Ctx, key, expiration)
	_, err := pipe.Exec(Ctx)
	if err != nil {
		return 0, err
	}
	return cmd.Result()
}

// LRange 获取列表指定范围内的元素
func LRange(key string, start, stop int64) ([]string, error) {
	return Rdb.LRange(Ctx, key, start, stop).Result()
}

// SAdd 向集合中添加一个或多个成员，不过期
func SAdd(key string, members ...interface{}) (int64, error) {
	return Rdb.SAdd(Ctx, key, members...).Result()
}

// SAddWithExp 向集合中添加一个或多个成员，带过期时间（秒）
func SAddWithExp(key string, expiration time.Duration, members ...interface{}) (int64, error) {
	pipe := Rdb.Pipeline()
	cmd := pipe.SAdd(Ctx, key, members...)
	pipe.Expire(Ctx, key, expiration)
	_, err := pipe.Exec(Ctx)
	if err != nil {
		return 0, err
	}
	return cmd.Result()
}

// SMembers 返回集合中的所有成员
func SMembers(key string) ([]string, error) {
	return Rdb.SMembers(Ctx, key).Result()
}

// ZAdd 向有序集合中添加一个或多个成员，或者更新已存在成员的分数，不过期
func ZAdd(key string, members ...*redis.Z) (int64, error) {
	return Rdb.ZAdd(Ctx, key, members...).Result()
}

// ZAddWithExp 向有序集合中添加一个或多个成员，或者更新已存在成员的分数，带过期时间（秒）
func ZAddWithExp(key string, expiration time.Duration, members ...*redis.Z) (int64, error) {
	pipe := Rdb.Pipeline()
	cmd := pipe.ZAdd(Ctx, key, members...)
	pipe.Expire(Ctx, key, expiration)
	_, err := pipe.Exec(Ctx)
	if err != nil {
		return 0, err
	}
	return cmd.Result()
}

// ZRangeByScore 根据分数范围获取有序集合中的成员
func ZRangeByScore(key string, opt *redis.ZRangeBy) ([]string, error) {
	return Rdb.ZRangeByScore(Ctx, key, opt).Result()
}

// 删除单个用户的缓存信息
func DelCacheUserById(key string) (err error) {
	err = Rdb.Del(Ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

// 订阅 channel
func SubChannel() {
	sub := Rdb.Subscribe(Ctx, "post_cache")
	ch := sub.Channel()
	go func() {
		for msg := range ch {
			if err := DeleteKey(msg.Payload); err != nil {
				log.Println("delete key ERROR:", err)
			}
		}
	}()
}

// 查询用户key是否存在
func ExistKey(key string) bool {
	n, err := Rdb.Exists(Ctx, key).Result()
	if err != nil {
		log.Println("find exist user Key error :", err)
	}
	if n == 0 {
		log.Println(key, "key no exist")
		return false
	}
	log.Println(key, "key exist")
	return true
}

// 设置key过期
func SetKeyExpired(key string) (err error) {
	err = Rdb.ExpireAt(Ctx, key, time.Now().Add(-10*time.Second)).Err()
	if err != nil {
		log.Println("Set Key Expired ERROR:", err)
		return err
	}
	return nil
}

// 删除某个key
func DeleteKey(key string) (err error) {
	if err = Rdb.Del(Ctx, key).Err(); err != nil {
		log.Println("Delete Key  ERROR:", err)
		return err
	}
	return nil
}

// 扣减redis库存
func DeductionOfInventory(key string) (int64, error) {
	script := `
	if tonumber(redis.call('get', KEYS[1])) > 0 then
    	redis.call('decr', KEYS[1])
    	return 1
	else
    	return 0
	end
`
	result, err := Rdb.Eval(Ctx, script, []string{key}).Result()
	return result.(int64), err
}
func ValidateOrderToken(key, orderToken string) (bool, error) {
	script := `
		if redis.call('get', KEYS[1]) == ARGV[1] then
			return redis.call('del', KEYS[1])
		else
			return 0
		end
	`
	result, err := Rdb.Eval(Ctx, script, []string{key}, orderToken).Result()
	if err != nil {
		return false, err
	}
	return result.(int64) == 1, nil
}
