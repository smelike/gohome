package main

import (
	"fmt"
	"time"
)

type Counter struct {
	count int
}

var mapChan = make(chan map[string]Counter, 1)

func main() {
	syncChan := make(chan struct{}, 2)
	go func() { // 接收操作
		for {
			if elem, ok := <-mapChan; ok { // 从通道中接收值
				counter := elem["count"]
				counter.count++
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan <- struct{}{}
	}()

	go func() { // 发送操作
		countMap := map[string]Counter{
			"count": Counter{},
		}
		for i := 0; i < 5; i++ {
			mapChan <- countMap // 发送值
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v. [sender]\n", countMap)
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}

/* 输出结果
The count map: map[count:{0}]. [sender]
The count map: map[count:{0}]. [sender]
The count map: map[count:{0}]. [sender]
The count map: map[count:{0}]. [sender]
The count map: map[count:{0}]. [sender]
Stopped. [receiver]
*/
