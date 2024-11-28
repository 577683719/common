package job_client

import (
	httpClient "ecs_cloud/common/util/http"
	"ecs_cloud/common/util/logger"
	"ecs_cloud/config"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var Client *JobService

type JobService struct {
	adminAddresses string
	username       string
	password       string
	loginCookie    map[string]string
}
type XxlJobInfo struct {
	ID                     int64  `json:"id"`
	JobGroup               int64  `json:"jobGroup"`
	JobDesc                string `json:"jobDesc"`
	AddTime                string `json:"-"`
	UpdateTime             string `json:"-"`
	Author                 string `json:"author"`
	AlarmEmail             string `json:"-"`
	ScheduleType           string `json:"scheduleType"`
	ScheduleConf           string `json:"scheduleConf"`
	MisfireStrategy        string `json:"misfireStrategy"`
	ExecutorRouteStrategy  string `json:"executorRouteStrategy"`
	ExecutorHandler        string `json:"executorHandler"`
	ExecutorParam          string `json:"executorParam"`
	ExecutorBlockStrategy  string `json:"executorBlockStrategy"`
	ExecutorTimeout        int    `json:"executorTimeout"`
	ExecutorFailRetryCount int    `json:"executorFailRetryCount"`
	GlueType               string `json:"glueType"`
	GlueSource             string `json:"-"`
	GlueRemark             string `json:"glueRemark"`
	ChildJobID             string `json:"childJobId"`
	TriggerStatus          int    `json:"triggerStatus"`
	TriggerLastTime        int64  `json:"-"`
	TriggerNextTime        int64  `json:"-"`
}

// Parse JSON response
var XxlJobGrouResponse struct {
	Data []*XxlJobGroup `json:"data"`
}
var XxlJobInfoResponse struct {
	Data []*XxlJobInfo `json:"data"`
}

// XxlJobGroup represents a job group structure
type XxlJobGroup struct {
	ID           int64     `json:"id"`
	AppName      string    `json:"appname"`
	Title        string    `json:"title"`
	AddressType  int       `json:"addressType"`
	AddressList  string    `json:"addressList"`
	UpdateTime   time.Time `json:"updateTime"`
	RegistryList []string  `json:"registryList"`
}

// GetRegistryList retrieves and returns the registry list
func (g *XxlJobGroup) GetRegistryList() []string {
	if g.AddressList != "" {
		g.RegistryList = strings.Split(g.AddressList, ",")
	}
	return g.RegistryList
}
func InitJob() {
	Client = NewJobLoginService()
	logger.LogrusObj.Infof("初始化:jobClinet成功")
}

func NewJobLoginService() *JobService {
	if config.Conf == nil {
		return &JobService{
			adminAddresses: "http://1.94.44.163:8082/xxl-job-admin",
			username:       "admin",
			password:       "123456",
			loginCookie:    make(map[string]string),
		}
	}
	return &JobService{
		adminAddresses: config.Conf.EcsJob.ServerAddr,
		username:       config.Conf.EcsJob.UserName,
		password:       config.Conf.EcsJob.Password,
		loginCookie:    make(map[string]string),
	}
}
func (service *JobService) StartJobInfo(infoId int64) []byte {
	body, _, _ := httpClient.Client.PostFormDataStrAndCookie(service.adminAddresses+"/jobinfo/start", map[string]string{
		"id": strconv.FormatInt(infoId, 10),
	}, service.getCookie())
	return body
}

// 获取jobinfo
func (service *JobService) StopJobInfo(infoId int64) map[string]interface{} {
	body, _, _ := httpClient.Client.PostFormDataStrAndCookie(service.adminAddresses+"/jobinfo/stop", map[string]string{
		"id": strconv.FormatInt(infoId, 10),
	}, service.getCookie())
	var res map[string]interface{}
	json.Unmarshal(body, &res)
	return res
}

// 获取jobinfo
func (service *JobService) RemoveJobInfo(infoId int64) map[string]interface{} {
	body, _, _ := httpClient.Client.PostFormDataStrAndCookie(service.adminAddresses+"/jobinfo/remove", map[string]string{
		"id": strconv.FormatInt(infoId, 10),
	}, service.getCookie())
	var res map[string]interface{}
	json.Unmarshal(body, &res)
	return res
}

// 获取jobinfo
func (service *JobService) GetJobInfo(jobGroupId int64, executorHandler, jobDesc, author string) ([]*XxlJobInfo, error) {
	xxlJobInfo := &XxlJobInfoResponse
	body, _, err := httpClient.Client.PostFormDataStrAndCookie(service.adminAddresses+"/jobinfo/pageList", map[string]string{
		"jobGroup":        fmt.Sprint(jobGroupId),
		"executorHandler": executorHandler,
		"jobDesc":         jobDesc,
		"triggerStatus":   "-1",
	}, service.getCookie())
	json.Unmarshal(body, &xxlJobInfo)
	return xxlJobInfo.Data, err
}

// 获取job组
func (service *JobService) GetJobGroup(appname string, title string) ([]*XxlJobGroup, error) {
	xxlJobGroup := &XxlJobGrouResponse
	body, _, err := httpClient.Client.PostFormDataStrAndCookie(service.adminAddresses+"/jobgroup/pageList", map[string]string{
		"appname": appname,
		"title":   title,
	}, service.getCookie())
	json.Unmarshal(body, &xxlJobGroup)
	return xxlJobGroup.Data, err
}
func (service *JobService) UpdateJobJnfo(info *XxlJobInfo) map[string]interface{} {
	marshal, _ := json.Marshal(info)
	// 定义一个 map 用于存储解码后的 JSON 数据
	var data map[string]interface{}

	// 将 JSON 字符串解码到 map 中
	err := json.Unmarshal([]byte(marshal), &data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}
	body, _, _ := httpClient.Client.PostFormDataAndCookie(service.adminAddresses+"/jobinfo/add", data, service.getCookie())
	var res map[string]interface{}
	err = json.Unmarshal(body, &res)
	return res
}
func (service *JobService) AddJobJnfo(info *XxlJobInfo) map[string]interface{} {
	marshal, _ := json.Marshal(info)
	// 定义一个 map 用于存储解码后的 JSON 数据
	var data map[string]interface{}

	// 将 JSON 字符串解码到 map 中
	err := json.Unmarshal([]byte(marshal), &data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}
	body, _, _ := httpClient.Client.PostFormDataAndCookie(service.adminAddresses+"/jobinfo/add", data, service.getCookie())
	var res map[string]interface{}
	err = json.Unmarshal(body, &res)
	return res
}

func (service *JobService) login() {
	_, resp := httpClient.Client.PostFormData(service.adminAddresses+"/login", map[string]string{
		"userName": service.username,
		"password": service.password,
	})
	cookies := resp.Cookies()
	var loginCookie *http.Cookie
	for _, cookie := range cookies {
		if cookie.Name == "XXL_JOB_LOGIN_IDENTITY" {
			loginCookie = cookie
			break
		}
	}
	if loginCookie == nil {
		panic("get xxl-job cookie error!")
	}
	service.loginCookie["XXL_JOB_LOGIN_IDENTITY"] = loginCookie.Value
}

func (service *JobService) getCookie() string {
	for i := 0; i < 3; i++ {
		cookieStr := service.loginCookie["XXL_JOB_LOGIN_IDENTITY"]
		if cookieStr != "" {
			return "XXL_JOB_LOGIN_IDENTITY=" + cookieStr
		}
		service.login()
	}
	panic("get xxl-job cookie error!")
}
