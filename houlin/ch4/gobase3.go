package main

import (
	"fmt"
	"time"
)

func main() {
	name := "Eric"
	go func() {
		fmt.Printf("Hello, %s!\n", name)
	}()
	time.Sleep(time.Millisecond)
	name = "Harry"
	// 要同时问候多个人，名单如下：names := []string{"Eric", "Robert", "Jim", "Mark"}
	// runtime.Gosched()
}
