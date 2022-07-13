package main

import (
	router "sulink/app/route"
	"sulink/common"
)

type Params struct {
	A int `json:"a"`
	B int `json:"b"`
}

type Result = int

type Result2 struct {
	C int `json:"c"`
}

func Init() {
	nacos := new(common.Nacos)
	nacos.Connect("192.168.0.2", 8848, "gin.json", "gin")

	router.Register("8200")

	/*nacos := &common.Nacos{}
	nacos.ServiceConfig.Host = 8848
	nacos.ServiceConfig.Port = 8848
	common.Config{}*/
}

func main() {
	Init()

	/*result := new(Result)
	c, _ := jsonrpc4go.NewClient("http", "192.168.0.142", "9504") // the protocol is http
	err := c.Call("/test/Add", Params{1, 6}, result, false)
	fmt.Println(err) // nil
	fmt.Println(*result) // 7


	result2 := new(Result)
	c2, _ := jsonrpc4go.NewClient("http", "192.168.0.142", "3233") // the protocol is http
	err2 := c2.Call("/TestService/Add", Params{1, 6}, result2, false)
	fmt.Println(err2) // nil
	fmt.Println(*result2) // 7*/
}
