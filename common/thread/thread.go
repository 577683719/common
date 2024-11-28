package thread

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"time"
)

var ThreadPool *ants.Pool

func init() {
	poolSize := 30         // 核心线程数
	queueCapacity := 1024  // 阻塞队列容量
	keepAliveSeconds := 30 // 线程空闲时间

	options := []ants.Option{
		ants.WithNonblocking(false), // 非阻塞提交任务
		ants.WithExpiryDuration(time.Duration(keepAliveSeconds) * time.Second), // 空闲线程的存活时间
		ants.WithPanicHandler(func(interface{}) {
			fmt.Println("任务发生了 panic")
		}),
		ants.WithPreAlloc(true),                  // 开启预分配
		ants.WithMaxBlockingTasks(queueCapacity), // 阻塞队列容量
	}

	// 创建线程池
	pool, err := ants.NewPool(poolSize, options...)
	if err != nil {
		fmt.Println("创建协程池失败:", err)
		return
	}
	ThreadPool = pool
}
