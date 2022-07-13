package services

import (
	"encoding/json"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
)

type Nacos struct {
	NacosServerConfig struct {
		Addr string //the nacos server Addr
		Port uint64 //the nacos server port
	}

	NamingClient naming_client.INamingClient
	ConfigClient config_client.IConfigClient

	DataId        string
	Group         string
	ServiceConfig struct {
		Port int    `json:"port"`
		Test string `json:"test"`
	}
}

func (nacos *Nacos) Init() {
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: nacos.NacosServerConfig.Addr,
			Port:   nacos.NacosServerConfig.Port,
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

	nacos.CreateNamingClient(clientConfig, serverConfigs)
	nacos.CreateConfigClient(clientConfig, serverConfigs)
}

// 创建服务发现客户端
func (nacos *Nacos) CreateNamingClient(clientConfig constant.ClientConfig, serverConfigs []constant.ServerConfig, ) {
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)

	if err != nil {
		log.Fatalf("初始化nacos创建服务发现客户端失败: %s", err.Error())
	}

	nacos.NamingClient = namingClient
	nacos.registerInstance()
}

//注册实例
func (nacos *Nacos) registerInstance() {
	_, err := nacos.NamingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "192.168.0.142",
		Port:        8081,
		ServiceName: "server_user",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "shanghai"},
		ClusterName: "cluster-a", // 默认值DEFAULT
		GroupName:   "group-a",   // 默认值DEFAULT_GROUP
	})

	if err != nil {
		log.Fatalf("nacos注册服务失败: %s", err.Error())
	}
}

// 创建动态配置客户端
func (nacos *Nacos) CreateConfigClient(clientConfig constant.ClientConfig, serverConfigs []constant.ServerConfig, ) {
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)

	if err != nil {
		log.Fatalf("初始化nacos创建动态配置客户端失败: %s", err.Error())
	}

	nacos.ConfigClient = configClient

	nacos.getConfig()
	nacos.onChange()
}

// 获取配置
func (nacos *Nacos) getConfig() {
	content, err := nacos.ConfigClient.GetConfig(vo.ConfigParam{
		DataId: nacos.DataId,
		Group:  nacos.Group,
	})

	if err != nil {
		log.Fatalf("获取%s配置失败: %s", nacos.DataId, err.Error())
	}

	err = json.Unmarshal([]byte(content), &nacos.ServiceConfig)
	if err != nil {
		log.Fatalf("解析%s配置失败: %s", nacos.DataId, err.Error())
	}
}

// 监听配置
func (nacos *Nacos) onChange() {
	err := nacos.ConfigClient.ListenConfig(vo.ConfigParam{
		DataId: nacos.DataId,
		Group:  nacos.Group,
		OnChange: func(namespace, group, dataId, data string) {
			err := json.Unmarshal([]byte(data), &nacos.ServiceConfig)
			if err != nil {
				log.Fatalf("获取%s配置失败: %s", nacos.DataId, err.Error())
			}
		},
	})

	if err != nil {
		log.Fatalf("监听%s配置失败: %s", nacos.DataId, err.Error())
	}
}
