package env_util

import (
	"github.com/577683719/common/common/util/strutil"
	"github.com/spf13/viper"
	"strings"
)

func init() {
	viper.AutomaticEnv()
}
func GetPortIndex() int {
	//如果是本地开发环境端口号固定，线上环境启动获取端口号
	prot := viper.GetInt("port")
	return prot
}
func GetOrderAddr() string {
	//如果是本地开发环境端口号固定，线上环境启动获取端口号
	addr := viper.GetString("orderAddr")
	return addr
}

// 获取设备服务地址
func GetDeviceAddr() string {
	//如果是本地开发环境端口号固定，线上环境启动获取端口号
	addr := viper.GetString("deviceAddr")
	return addr
}

// 获取设备服务地址
func GetProductAddr() string {
	//如果是本地开发环境端口号固定，线上环境启动获取端口号
	addr := viper.GetString("productAddr")
	return addr
}

// 获取环境ipAddr
func GetNacos() string {
	env := viper.GetString("NACOS")
	return env
}

// 获取环境ipAddr
func GetNacosIpAddr() string {
	env := viper.GetString("IP_ADDR")
	return env
}

// 获取环境port
func GetNacosPort() string {
	env := viper.GetString("PORT")
	return env
}

// 获取环境namespaceID
func GetNacosNamespaceID() string {
	env := viper.GetString("NAMESPACE_ID")
	return env
}

// 获取环境group
func GetNacosGroup() string {
	env := viper.GetString("GROUP")
	return env
}

// 获取环境dataID
func GetNacosDataID() string {
	env := viper.GetString("DATA_ID")
	return env
}

// 获取环境format
func GetNacosFormat() string {
	env := viper.GetString("FORMAT")
	return env
}

// 获取环境logDir
func GetNacosLogDir() string {
	env := viper.GetString("LOG_DIR")
	return env
}

// 获取环境cacheDir
func GetNacosCacheDir() string {
	env := viper.GetString("CACHE_DIR")
	return env
}

// 获取环境logLevel
func GetNacosLogLevel() string {
	env := viper.GetString("LOG_LEVEL")
	return env
}

// 获取环境
func GetEnv() string {
	env := viper.GetString("env")
	return env
}

// 生产环境
func IsProd() bool {
	if GetEnv() == "prod" {
		return true
	}
	return false
}

// 开发环境
func IsDev() bool {
	if GetEnv() == "dev" || strutil.IsEmpty(GetEnv()) {
		return true
	}
	return false
}

// 开发环境
func IsTest() bool {
	if GetEnv() == "test" {
		return true
	}
	return false
}

// 获取账号服务地址
func GetAccountAddr() string {
	//如果是本地开发环境端口号固定，线上环境启动获取端口号
	addr := viper.GetString("accountAddr")
	return addr
}

func GetIp() string {
	ip := viper.GetString("IP")
	return ip
}
func GetPort() string {
	port := viper.GetString("port")
	return port
}

func GetActivityAddr() string {
	//如果是本地开发环境端口号固定，线上环境启动获取端口号
	addr := viper.GetString("activityAddr")
	return addr
}

// 获取消息服务地址
func GetMessageAddr() string {
	//如果是本地开发环境端口号固定，线上环境启动获取端口号
	addr := viper.GetString("messageAddr")
	return addr
}

// 获取文件服务地址多个
func GetFilesHttpAddrs() []string {
	//如果是本地开发环境端口号固定，线上环境启动获取端口号
	addrs := viper.GetString("filesHttpAddrs")
	filesHttpAddrs := strings.Split(addrs, ",")
	if len(filesHttpAddrs) == 1 {
		if strutil.IsEmpty(filesHttpAddrs[0]) {
			return nil
		}
	}
	return filesHttpAddrs
}

// 获取文件服务地址多个
func GetFileServicePort() string {
	//如果是本地开发环境端口号固定，线上环境启动获取端口号
	filePort := viper.GetString("fileServicePort")

	return filePort
}
