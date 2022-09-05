package main

import (
	"fmt"
	"time"
)

/*
	缓冲通道是以异步的方式传递元素值；
	非缓冲通道只能同步地传递元素值。
*/

func main() {
	sendingInterval := time.Second
	receptionInterval := time.Second * 2
	intChan := make(chan int, 0) // 非缓冲通道
	// intChan := make(chan int, 5) // capSize =5 缓冲通道

	// 发送操作
	go func() {
		var ts0, ts1 int64
		for i := 1; i <= 5; i++ {
			intChan <- i
			ts1 = time.Now().Unix()
			if ts0 == 0 {
				fmt.Println("Sent:", i)
			} else {
				fmt.Printf("Sent: %d [interval: %d s]\n", i, ts1-ts0)
			}
			ts0 = time.Now().Unix()
			time.Sleep(sendingInterval)
		}
		close(intChan)
	}()

	// 接收操作
	var ts0, ts1 int64
Loop:
	for {
		select {
		case v, ok := <-intChan:
			if !ok {
				break Loop
			}
			ts1 = time.Now().Unix()
			if ts0 == 0 {
				fmt.Println("Received:", v)
			} else {
				fmt.Printf("Received: %d [interval: %d s]\n", v, ts1-ts0)
			}
		}
		ts0 = time.Now().Unix()
		time.Sleep(receptionInterval)
	}
	fmt.Println("End.")
}

/*

// happen before

缓冲通道
Sent: 1
Received: 1
Sent: 2 [interval: 2 s]
Received: 2 [interval: 2 s]
Received: 3 [interval: 2 s]
Sent: 3 [interval: 2 s]
Received: 4 [interval: 2 s]
Sent: 4 [interval: 2 s]
Received: 5 [interval: 2 s]
Sent: 5 [interval: 2 s]
End.

// 非缓冲通道
Sent: 1
Received: 1
Received: 2 [interval: 2 s]
Sent: 2 [interval: 2 s]
Received: 3 [interval: 2 s]
Sent: 3 [interval: 2 s]
Received: 4 [interval: 2 s]
Sent: 4 [interval: 2 s]
Received: 5 [interval: 2 s]
Sent: 5 [interval: 2 s]
End.

*/
