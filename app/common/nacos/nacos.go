package nacos

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type Nacos struct {
	namingClient naming_client.INamingClient
	configClient config_client.IConfigClient
	dataId       string
	group        string
	Config       Config
}

type Config struct {
	Port int    `json:"port"`
	Test string `json:"test"`
}

var once sync.Once
var instance *Nacos

func GetNacosInstance() *Nacos {
	once.Do(func() {
		instance = new(Nacos)
	})
	return instance
}

// 建立连接
func (nacos *Nacos) Connect(Host string, Port uint64, DataId string, Group string) {
	nacos.dataId = DataId
	nacos.group = Group

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: Host,
			Port:   Port,
		},
	}

	// 创建clientConfig
	clientConfig := constant.ClientConfig{
		// 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		NamespaceId:         "44ee8d7a-d34c-488a-af5d-41104dbafcda",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		CacheDir:            "runtime/cache/",
		LogDir:              "runtime/log/",
		LogLevel:            "debug",
	}

	// 创建服务发现客户端
	nacos.CreateNamingClient(clientConfig, serverConfigs)

	// 创建动态配置客户端
	nacos.CreateConfigClient(clientConfig, serverConfigs)
}

// 创建服务发现客户端
func (nacos *Nacos) CreateNamingClient(clientConfig constant.ClientConfig, serverConfigs []constant.ServerConfig) {
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)

	if err != nil {
		log.Fatalf("初始化nacos创建服务发现客户端失败: %s", err.Error())
	}

	nacos.namingClient = namingClient
	nacos.registerInstance()
}

//注册实例
func (nacos *Nacos) registerInstance() {
	_, err := nacos.namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "192.168.0.171",
		Port:        8200,
		ServiceName: "server_business",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "shanghai"},
		ClusterName: "cluster-sulink", // 默认值DEFAULT
		GroupName:   "sulink",         // 默认值DEFAULT_GROUP
	})

	if err != nil {
		log.Fatalf("nacos注册服务失败: %s", err.Error())
	}
}

// 创建动态配置客户端
func (nacos *Nacos) CreateConfigClient(clientConfig constant.ClientConfig, serverConfigs []constant.ServerConfig) {
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)

	if err != nil {
		log.Fatalf("初始化nacos创建动态配置客户端失败: %s", err.Error())
	}

	nacos.configClient = configClient

	nacos.getConfig()
	nacos.onChange()
}

// 获取配置
func (nacos *Nacos) getConfig() {
	content, err := nacos.configClient.GetConfig(vo.ConfigParam{
		DataId: nacos.dataId,
		Group:  nacos.group,
	})

	if err != nil {
		log.Fatalf("获取%s配置失败: %s", nacos.dataId, err.Error())
	}

	err = json.Unmarshal([]byte(content), &nacos.Config)
	if err != nil {
		log.Fatalf("解析%s配置失败: %s", nacos.dataId, err.Error())
	}
}

// 监听配置
func (nacos *Nacos) onChange() {
	err := nacos.configClient.ListenConfig(vo.ConfigParam{
		DataId: nacos.dataId,
		Group:  nacos.group,
		OnChange: func(namespace, group, dataId, data string) {
			err := json.Unmarshal([]byte(data), &nacos.Config)
			if err != nil {
				log.Fatalf("获取%s配置失败: %s", nacos.dataId, err.Error())
			}
			log.Println(nacos.Config)
		},
	})

	if err != nil {
		log.Fatalf("监听%s配置失败: %s", nacos.dataId, err.Error())
	}
}

func (nacos Nacos) SelectOneHealthyInstance(serviceName string) (ip string, port uint64) {
	instance, err := nacos.namingClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: serviceName,
		GroupName:   "sulink",                   // 默认值DEFAULT_GROUP
		Clusters:    []string{"cluster-sulink"}, // 默认值DEFAULT
	})
	if err != nil {
		log.Fatal("SelectOneHealthyInstance error " + err.Error())
	}
	return instance.Ip, instance.Port
}
