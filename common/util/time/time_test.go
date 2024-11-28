package time_util

import (
	"fmt"
	"github.com/577683719/common/common/util/common_util/date_util"
	"github.com/577683719/common/common/util/logger"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	ZeroPointExecution(func() {
		logger.LogrusObj.Errorf("每天1凌晨执行:%s", date_util.DateToDateTimeStr(time.Now()))
	})
	tomorrow := CalculateDurationUntilTomorrow()
	fmt.Println(tomorrow.Seconds())
	time.Sleep(1000 * time.Second)
}
func TestTime2(t *testing.T) {
	ZeroPointExecution(func() {
		fmt.Printf("每天2凌晨执行:%s", date_util.DateToDateTimeStr(time.Now()))
	})
	tomorrow := CalculateDurationUntilTomorrow()
	fmt.Println(tomorrow.Seconds())
	time.Sleep(1000 * time.Second)
}
