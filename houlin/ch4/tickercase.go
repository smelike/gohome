package main

import (
	"fmt"
	"time"
)

/*
	思考：为什么总是打印不出 End. [sender] ?
	使用了 ticker.Stop() 还是不出现 End. [sender]
*/

func main() {
	intChan := make(chan int, 1)

	/*
		// 断续器，一旦被初始化，所有的绝对到期时间就已确定。
		// 固定不变的到期时间，使得断续器非常适用于定时任务的触发器。
		// *time.Ticker 类型的方法集合中只有一个方法——Stop，停止断续器。
		// 一旦断续器停止，就不会再向其通知通道发送任何元素值了。
		如果此时字段 C 中已经有了一个元素值，那么该元素值会一直在那里，直至被接收。

	*/
	ticker := time.NewTicker(time.Second)
	go func() {
		for _ = range ticker.C { // ticker.C 定时触发器
			select { // 随机算法处理 case
			case intChan <- 1: // 给 intChan 通道发送 1
			case intChan <- 2: // 给 intChan 通道发送 2
			case intChan <- 3: // 给 intChan 通道发送 3
			}
		}
		fmt.Println("End. [sender]")
	}()
	var sum int
	for e := range intChan { // 接收数据
		fmt.Printf("Received: %v\n", e)
		sum += e
		if sum > 10 {
			fmt.Printf("Got: %v\n", sum)
			ticker.Stop() // 发现没有 End. [sender] 后添加的
			// time.Sleep(time.Second)
			break
		}
	}
	// ticker.Stop() // 发现没有 End. [sender] 后添加的
	fmt.Println("End. [receiver]")
}
