package sentinel_util

import (
	"ecs_cloud/common/util/logger"
	"ecs_cloud/config"
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	sentinel_config "github.com/alibaba/sentinel-golang/core/config"
	"github.com/alibaba/sentinel-golang/ext/datasource"
	"github.com/alibaba/sentinel-golang/logging"
	"github.com/alibaba/sentinel-golang/pkg/datasource/nacos"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
)

const flowDataIdPostfix = "-sentinel-flow" // 流控规则配置文件后缀
const groupId = "flow"                     // 配置文件所属分组

// 初始化Nacos
func InitSentinel() {
	// We should initialize Sentinel first.
	conf := sentinel_config.NewDefaultConfig()
	// for testing, logging output to console
	conf.Sentinel.Log.Logger = logging.NewConsoleLogger()
	err := sentinel.InitWithConfig(conf)
	if err != nil {
		logger.LogrusObj.Error(err)
	}
	serverConfig := []constant.ServerConfig{
		{
			ContextPath: "/nacos",
			IpAddr:      config.NacosConf.Nacos.IpAddr, //nacos 地址
			Port:        config.NacosConf.Nacos.Port,   //nacos 端口
			Scheme:      "http",
		},
	}

	clientConfig := &constant.ClientConfig{
		NamespaceId: config.NacosConf.Nacos.NamespaceID, //命名空间 比较重要 拿取刚才创建的命名空间ID
		TimeoutMs:   5000,
	}
	// 生成nacos config client(配置中心客户端)
	client, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfig,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		fmt.Printf("Fail to create client, err: %+v", err)
		return
	}

	// 注册流控规则Handler
	h := datasource.NewFlowRulesHandler(datasource.FlowRuleJsonArrayParser)

	// 创建NacosDataSource数据源
	// groupId 对应在nacos中创建配置文件的group
	// dataId 对应在nacos中创建配置文件的dataId，一定要和Nacos配置中心中的配置文件名一致，否则无法匹配
	dataId := config.Conf.Server.Name + flowDataIdPostfix
	nds, err := nacos.NewNacosDataSource(client, groupId, dataId, h)
	if err != nil {
		fmt.Printf("Fail to create nacos data source client, err: %+v", err)
		return
	}

	//nacos数据源初始化
	err = nds.Initialize()
	logger.LogrusObj.Infof("流控规则设置成功")
	if err != nil {
		fmt.Printf("Fail to initialize nacos data source client, err: %+v", err)
		return
	}
}
