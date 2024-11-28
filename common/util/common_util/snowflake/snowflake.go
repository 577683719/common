package common_util_snowflake

import (
	"fmt"
	"github.com/duke-git/lancet/v2/datetime"
	"math/rand"
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func init() {
	var st time.Time
	var err error
	st, err = time.Parse("2006-01-02", datetime.FormatTimeToStr(time.Now(), "yyyy-MM-dd"))
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	// 使用当前时间的Unix时间戳作为随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生成 10 到 200 范围内的随机数
	randomNumber := rand.Intn(1000-10+1) + 10

	nodeNumber := int64(randomNumber)
	node, err = sf.NewNode(nodeNumber)
	fmt.Println(fmt.Sprintf("当前node是 %d", nodeNumber))
}

func GenID() int64 {
	return node.Generate().Int64()
}
func GenIDStr() string {
	return node.Generate().String()
}
