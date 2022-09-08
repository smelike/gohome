package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 1)

	// 发送操作
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			intChan <- i
		}
		close(intChan)
	}()
	// 定时器-超时设计
	timeout := time.Millisecond * 500
	var timer *time.Timer
	for {
		if timer == nil {
			timer = time.NewTimer(timeout) // timer 为 nil 创建定时器，
		} else {
			timer.Reset(timeout)
		}
		select { // select
		case e, ok := <-intChan:
			if !ok {
				fmt.Println("End.")
				return
			}
			fmt.Printf("Received: %v\n", e)
		case <-timer.C:
			fmt.Println("Timeout!")
		}
	}
}
