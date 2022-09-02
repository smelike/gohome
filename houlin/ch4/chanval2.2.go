package main

import (
	"fmt"
	"time"
)

// 结构体声明
type Counter struct {
	count int
}

// 结构体 Counter 的方法
func (counter *Counter) String() string {
	return fmt.Sprintf("{counter: %d}", counter.count)
}

// 全局通道声明，结构体 Counter 的指针(*Counter)作为元素类型
var mapChan = make(chan map[string]*Counter, 1)

func main() {

	// 同步信号的通道，队列大小为 2
	syncChan := make(chan struct{}, 2)
	// goroutine1 - channel - 接收操作
	go func() {
		for {
			if elem, ok := <-mapChan; ok {
				counter := elem["count"]
				counter.count++
			} else {
				break
			}
		}
		// 打印输出停止接收提示
		fmt.Println("Stopped. [receiver]")
		// 给同步信号的通道发送值
		syncChan <- struct{}{}
	}()

	// * 指针的意思
	// & 获取指针，即内存地址
	// goroutine2 - channel - 发送操作
	go func() {
		countMap := map[string]*Counter{
			"count": &Counter{}, // & 获取指针
		}
		for i := 0; i < 5; i++ {
			mapChan <- countMap // 发送countMap 给通道 mapChan
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v. [sender]\n", countMap)
		}
		close(mapChan)         // 关闭通道发送操作
		syncChan <- struct{}{} // 给同步信号通道发送
	}()
	<-syncChan // 通道取值操作，但通道中值为 nil 时，该操作将使得当前的 goroutine 阻塞
	<-syncChan
}

/*
	The count map: map[count:0xc000012088]. [sender]
The count map: map[count:0xc000012088]. [sender]
The count map: map[count:0xc000012088]. [sender]
The count map: map[count:0xc000012088]. [sender]
The count map: map[count:0xc000012088]. [sender]
Stopped. [receiver]
*/
