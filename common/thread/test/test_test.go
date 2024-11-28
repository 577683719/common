package test

import (
	"fmt"
	"github.com/577683719/common/common/thread"
	"testing"
	"time"
)

func TestWorkerPool(t *testing.T) {
	// 自定义拒绝策略

	defer thread.ThreadPool.Release()
	// 提交任务到线程池
	taskCount := 10
	for i := 0; i < taskCount; i++ {
		taskIndex := i
		err := thread.ThreadPool.Submit(func() {
			fmt.Printf("开始执行任务 %d\n", taskIndex)
			time.Sleep(1 * time.Second) // 模拟任务执行时间
			fmt.Printf("任务 %d 执行完成\n", taskIndex)
		})
		if err != nil {
			fmt.Printf("提交任务 %d 失败: %v\n", taskIndex, err)
		}
		if i%10 == 0 {
			//time.Sleep(1 * time.Second)
		}
	}

	// 等待所有任务完成
	time.Sleep(100 * time.Minute)

	fmt.Println("所有任务已完成")
}
