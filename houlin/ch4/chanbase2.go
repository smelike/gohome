package main

import (
	"fmt"
	"time"
)

var strChan = make(chan string, 3)

func main() {
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)

	go receive(strChan, syncChan1, syncChan2) // 接收操作
	go send(strChan, syncChan1, syncChan2)    // 发送操作
	<-syncChan2
	<-syncChan2
}

// 通过 syncChan1 做信号控制，实现对 synChan 中数据的接收
func receive(strChan <-chan string,
	syncChan1 <-chan struct{}, // 接收
	syncChan2 chan<- struct{}) { // 发送
	<-syncChan1
	fmt.Println("Received a sync signal and wait a second... [receiver]")
	time.Sleep(time.Second)
	for {
		if elem, ok := <-strChan; ok {
			fmt.Println("Received:", elem, "[receiver]")
		} else {
			break
		}
	}
	fmt.Println("Stopped. [receiver]")
	syncChan2 <- struct{}{}
}

func send(strChan chan<- string,
	syncChan1 chan<- struct{},
	syncChan2 chan<- struct{}) {
	for _, elem := range []string{"a", "b", "c", "d"} {
		strChan <- elem
		fmt.Println("Sent:", elem, "[sender]")
		if elem == "c" {
			syncChan1 <- struct{}{}
			fmt.Println("Sent a sync signal. [sender]")
		}
	}
	fmt.Println("Wait 2 seconds... [sender]")
	time.Sleep(time.Second * 2)
	close(strChan)
	syncChan2 <- struct{}{}
}

/*
	运行输出结果：

	PS D:\gohome\houlin\ch4> go run .\chanbase2.go
Sent: a [sender]
Sent: b [sender]
Sent: c [sender]
Sent a sync signal. [sender]
Received a sync signal and wait a second... [receiver]
Sent: d [sender]
Wait 2 seconds... [sender]
Received: a [receiver]
Received: b [receiver]
Received: c [receiver]
Received: d [receiver]
Stopped. [receiver]
*/
