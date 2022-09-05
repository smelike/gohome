package main

import (
	"fmt"
	"time"
)

// 定时器的字段 C，可以及时得到定时器到期的通知
// 创建一个定时器 timer := time.NewTimer(time.Second)
func main() {
	intChan := make(chan int, 1)
	go func() {
		time.Sleep(time.Second)
		intChan <- 1
	}()
	select {
	case e := <-intChan:
		fmt.Printf("Received: %v\n", e)
	case <-time.NewTimer(time.Millisecond * 500).C:
		// time.After(time.Millisecond * 500)
		fmt.Println("Timeout!")
	}
}
