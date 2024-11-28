package time_util

import (
	"fmt"
	"github.com/duke-git/lancet/v2/datetime"
	"time"
)

func ZeroPointExecution(invoke func()) {
	ticker := time.NewTicker(CalculateDurationUntilTomorrow())
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("每天凌晨的定时器触发-调用开始")
				invoke()
				fmt.Println("每天凌晨的定时器触发-调用结束")
				//重新计时
				tomorrow := CalculateDurationUntilTomorrow()
				ticker.Reset(tomorrow)
				fmt.Println("每天凌晨的定时器触发-重新计时")
			}
		}
	}()

}

func CalculateDurationUntilTomorrow() time.Duration {
	// 获取当前时间
	now := time.Now()
	return datetime.BeginOfDay(datetime.AddDay(time.Now(), 1)).Sub(now)
}
