package gredis_lock

import (
	"ecs_cloud/common/util/gredis"
	"github.com/go-redsync/redsync/v4"
	"time"
)

func GetLockTimeOut(logckKey string, lockTimeout time.Duration) *redsync.Mutex {
	// 设置锁的超时时间为 10 秒

	// 创建锁对象，并指定超时时间
	mutex := gredis.Redsync.NewMutex(
		logckKey,
		redsync.WithExpiry(lockTimeout),
		redsync.WithRetryDelay(100*time.Millisecond), // 设置重试延迟为 100 毫秒
		redsync.WithTries(50),                        //重试50次
	)
	return mutex
}

// NewLock 实例化一个分布式锁，用来实现幂等，降低重试成本
func GetLock(logckKey string) *redsync.Mutex {
	// 设置锁的超时时间为 10 秒
	lockTimeout := 10 * time.Second

	// 创建锁对象，并指定超时时间
	mutex := gredis.Redsync.NewMutex(
		logckKey,
		redsync.WithExpiry(lockTimeout), //设置锁的过期时间为 10 秒
		redsync.WithRetryDelay(100*time.Millisecond), // 设置重试延迟为 100 毫秒
		redsync.WithTries(50),                        //重试50次
	)
	return mutex
}

// LockRelet 周期性续租,过去无可挽回，未来可以改变
// num定义时间：单位毫秒
// size定义续租的次数
func LockRelet(num int, size int, mutex *redsync.Mutex) chan bool {
	done := make(chan bool)
	if size <= 0 {
		return nil
	}
	go func() {
		ticker := time.NewTicker(time.Duration(num) * time.Millisecond)
		defer ticker.Stop()
		for size > 0 {
			size--
			select {
			case <-ticker.C:
				extend, err := mutex.Extend()
				if err != nil {
				} else if !extend {
				}
			case <-done:
				return
			}
		}
	}()
	return done
}
