package main

import (
	"fmt"

	"github.com/sunquakes/jsonrpc4go"
)

type Params struct {
	A int `json:"a"`
	B int `json:"b"`
}

type Result = int

type Result2 struct {
	C int `json:"c"`
}

func main() {
	result := new(Result)
	c, _ := jsonrpc4go.NewClient("http", "192.168.0.97", "9504") // the protocol is http
	err := c.Call("/addition/add", Params{100, 600}, result, false)
	fmt.Println(err)     // nil
	fmt.Println(*result) // 7

	/*
		result2 := new(Result)
		c2, _ := jsonrpc4go.NewClient("http", "192.168.0.142", "3233") // the protocol is http
		err2 := c2.Call("/TestService/Add", Params{1, 6}, result2, false)
		fmt.Println(err2) // nil
		fmt.Println(*result2) // 7 */
}
