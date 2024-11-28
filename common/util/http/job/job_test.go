package job_client

import (
	"ecs_cloud/common/util/common_util/corn_util"
	"fmt"
	"github.com/duke-git/lancet/v2/compare"
	"github.com/duke-git/lancet/v2/datetime"
	"strconv"
	"testing"
	"time"
)

func TestJob(t *testing.T) {
	InitJob()
	//		"appname": "golang-jobs",
	//		"title":   "golang执行器",
	group, _ := Client.GetJobGroup("golang-jobs", "按日扣费释放")
	fmt.Println(group[0].ID)
	//info := Client.GetJobInfo(group[0].ID, "task.auto")
	//fmt.Println(info[0].ID)
	//使用定时器
	fmt.Println(group[0].ID)
	handler := "task.ReleasedByDayDeductions"
	cons := corn_util.DateTimeToCronStr(datetime.AddMinute(time.Now(), 1))
	//实例id

	//注释
	for i := 0; i < 1; i++ {
		instanceId := strconv.Itoa(i)
		xxlJobInfo := funcName(cons, group, handler, instanceId)

		println(fmt.Sprint(xxlJobInfo))
		jnfo := Client.AddJobJnfo(xxlJobInfo)
		fmt.Println(jnfo)
		fmt.Println(compare.EqualValue(jnfo["code"], 200))
	}

	info, _ := Client.GetJobInfo(group[0].ID, handler, "实例:1", "用户:1")
	println(info[0].ID)
	//xxlJobInfo.ID = info[0].ID
	//xxlJobInfo.Author = "李昕宇---"
	//xxlJobInfo.ScheduleConf = common_util.AddCornDay(info[0].ScheduleConf, 10)
	//jnfo := Client.UpdateJobJnfo(xxlJobInfo)
	//fmt.Printf("更新成功:%s", jnfo)

	//remove := Client.RemoveJobInfo(info[0].ID)
	//remove1 := Client.RemoveJobInfo(info[1].ID)
	//fmt.Println(remove)
	//fmt.Println(remove1)

}

func funcName(cons string, group []*XxlJobGroup, handler string, instanceId string) *XxlJobInfo {
	var desc = fmt.Sprintf("用户:%v", 1)
	var desc1 = fmt.Sprintf("实例:%v", 1)
	xxlJobInfo := &XxlJobInfo{
		JobGroup:               group[0].ID,
		JobDesc:                desc1,
		Author:                 desc,
		ScheduleType:           "CRON",
		ScheduleConf:           cons,
		MisfireStrategy:        "FIRE_ONCE_NOW",
		ExecutorRouteStrategy:  "ROUND",
		ExecutorHandler:        handler,
		ExecutorParam:          instanceId,
		ExecutorBlockStrategy:  "SERIAL_EXECUTION",
		ExecutorTimeout:        30,
		ExecutorFailRetryCount: 3,
		GlueType:               "BEAN",
		GlueRemark:             "",
		TriggerStatus:          1,
	}
	return xxlJobInfo
}
