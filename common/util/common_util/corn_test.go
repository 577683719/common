package common_util

import (
	"fmt"
	"github.com/577683719/common/common/util/common_util/corn_util"
	"github.com/duke-git/lancet/v2/datetime"
	"testing"
	"time"
)

func TestName1(t *testing.T) {
	for i := 0; i < 100; i++ {
		println(GenerateUniqueID("SID", 8))
	}
}
func TestName(t *testing.T) {

	// 测试数据
	data := make([]int, 250)
	for i := 0; i < 250; i++ {
		data[i] = i + 1
	}

	// 将数组按每 100 个元素进行分组
	chunks := GroupList(data, 100)

	// 输出结果
	for i, chunk := range chunks {
		fmt.Printf("Group %d: %v\n", i+1, chunk)
	}

	//for i := 0; i < 100; i++ {
	//	xid_util.New()
	//println(fmt.Sprintf("%s%s%d", date_util.DateTimeToDateTimeNumberStr(time.Now()), GetRandomCode(2), time.Now().UnixMilli()))
	//unix := time.Now().UnixMicro()
	////unix1 := time.Now().UnixNano()
	//unix2 := time.Now().UnixMilli()
	//println(unix)
	////println(unix1)
	//println(unix2)
	//next := xid_util.Next()
	//println(next)
	//println(next)
	//}
}
func TestAddCornDate(t *testing.T) {
	// 定义一个 Cron 表达式
	expression := corn_util.DateTimeToCronStr(datetime.AddMinute(time.Now(), 1))
	fmt.Println(expression)
	day := corn_util.AddCornDay(expression, 1)

	fmt.Println(day)
	fmt.Println(corn_util.AddCornMinute(expression, 1))
	fmt.Println(corn_util.AddCornWeek(expression, 1))
	fmt.Println(corn_util.AddCornMonth(expression, 1))
}
func TestGetRandom(t *testing.T) {
	number := RandomNumber(100000, 110000)
	fmt.Println(number)

}
