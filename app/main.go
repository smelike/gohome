package main

import (
	router "sulink/app/route"
	"sulink/common/jsonrpc"
	"sulink/common/nacos"
)

func Init() {
	nacos := nacos.GetNacosInstance()
	nacos.Connect("192.168.0.2", 8848, "gin.json", "gin")

	go func() {
		jsonrpc := jsonrpc.GetJsonRpcInstance()
		jsonrpc.Register("192.168.0.110", "8200")
	}()

	router.Register("8100")
}

func main() {
	Init()
}
