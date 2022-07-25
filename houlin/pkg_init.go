package main

import (
	"fmt"
	"runtime"
)

func init() {
	fmt.Printf("Map: %v\n", m)
	info = fmt.Sprintf("OS: %s, Arch: %s", runtime.GOOS, runtime.GOARCH)
}

// 全局变量的初始化都会在代码包的初始化函数执行前完成
// 全局变量的初始化
var m = map[int]string{1: "A", 2: "B", 3: "C"}

// 全局变量的初始化
var info string

func main() {
	fmt.Printf(info)
}
