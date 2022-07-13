package jsonrpc

import (
	"log"
	"strconv"
	"sulink/app/services"
	"sulink/common/nacos"
	"sync"

	"github.com/sunquakes/jsonrpc4go"
)

type JsonRpc struct {
}

var once sync.Once
var instance *JsonRpc

func GetJsonRpcInstance() *JsonRpc {
	once.Do(func() {
		instance = new(JsonRpc)
	})
	return instance
}

//注册服务
func (that JsonRpc) Register(ip string, port string) {
	s, _ := jsonrpc4go.NewServer("http", ip, port) // the protocol is http
	s.Register(new(services.BusinessService))
	s.Start()
}

//服务调度
func (that JsonRpc) Call(serviceName string, method string, params map[string]interface{}) map[string]interface{} {
	nacos := nacos.GetNacosInstance()
	ip, port := nacos.SelectOneHealthyInstance(serviceName)
	var result map[string]interface{}
	cc, _ := jsonrpc4go.NewClient("http", ip, strconv.Itoa(int(port))) // the protocol is http
	err := cc.Call("/TestService/Add", params, &result, false)

	if err != nil {
		log.Println("jsonrpc call serviceName=" + serviceName + ", method=" + method + " error")
	}

	return result
}
