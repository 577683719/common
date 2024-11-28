package device_client

import (
	httpClient "github.com/577683719/common/common/util/http"
	"github.com/577683719/common/config"
)

import (
	"encoding/json"
	"fmt"
)

type InstanceInfo struct {
	CpuType      string `json:"CpuType"`
	CpuNum       int    `json:"CpuNum"`
	CpuMem       string `json:"CpuMem"`
	CpuFreq      string `json:"CpuFreq"`
	SysDiskSize  string `json:"SysDiskSize"`
	DataDiskSize string `json:"DataDiskSize"`
	GpuType      string `json:"GpuType"`
	GpuNum       int    `json:"GpuNum"`
	GpuMem       string `json:"GpuMem"`
	CudaVer      string `json:"CudaVer"`
	Image        string `json:"Image"`
}
type InstanceCreateReqeust struct {
	InstanceInfo InstanceInfo `json:"InstanceInfo"`
	InstanceID   string       `json:"InstanceID"`
	UserID       int64        `json:"UserID"`
	BmxSrvArea   string       `json:"BmxSrvArea"`
	ProductID    string       `json:"ProductID"`
	Priority     int          `json:"Priority"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	InstanceID     string `json:"InstanceID"`
	InstanceStatus string `json:"InstanceStatus"`
}
type Request struct {
	InstanceID string `json:"InstanceID"`
}

const INSTANCE_ALLOCATE = "/instance-allocate"
const INSTANCE_RELEASE = "/instance-release"
const POWER_OFFURL = "/instance-poweroff"
const INSTANCE_POWERON = "/instance-poweron"
const GET_INSTANCEID = "/instanceid-generate"

// 创建实例
func InstanceCreate(instance InstanceCreateReqeust, res *Response) *Response {
	body := httpClient.Client.PostJSON(config.Conf.DeviceMgr.Address+INSTANCE_ALLOCATE, instance)
	json.Unmarshal(body, &res)
	data, _ := json.Marshal(res)
	fmt.Println(string(data))
	return res
}

// 实例释放
func InstanceRelease(req Request, res *Response) *Response {
	body := httpClient.Client.PostJSON(config.Conf.DeviceMgr.Address+INSTANCE_RELEASE, req)
	json.Unmarshal(body, &res)
	data, _ := json.Marshal(res)
	fmt.Println(string(data))
	return res
}

// 实例关机
func InstancePowerOff(req Request, res *Response) *Response {
	body := httpClient.Client.PostJSON(config.Conf.DeviceMgr.Address+POWER_OFFURL, req)
	json.Unmarshal(body, &res)
	data, _ := json.Marshal(res)
	fmt.Println(string(data))
	return res
}

// 实例开机
func InstancePowerOn(req Request, res *Response) *Response {
	body := httpClient.Client.PostJSON(config.Conf.DeviceMgr.Address+INSTANCE_POWERON, req)
	json.Unmarshal(body, &res)
	data, _ := json.Marshal(res)
	fmt.Println(string(data))
	return res
}

// 实例开机
func GetInstanceId(res *Response) *Response {
	body := httpClient.Client.PostJSON(config.Conf.DeviceMgr.Address+GET_INSTANCEID, "")
	json.Unmarshal(body, &res)
	data, _ := json.Marshal(res)
	fmt.Println(string(data))
	return res
}
