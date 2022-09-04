package main

import (
	"fmt"
	"time"
)

/*
	从一个还未被初始化的通道中接收元素值会导致当前 goroutine 的永久阻塞。
	当通道中没有任何元素值时，for 语句所在的 goroutine 也会陷入阻塞，阻塞的具体位置会在其中的 range 子句处。

	for 语句会不断地尝试从通道中接收元素值，直到该通道关闭。在通道关闭时，如果通道中已无元素值，那么这条 for 语句的执行就会立即结束。
	而当此时的通道中还有遗留的元素值时，for 语句仍可以继续把它们去完。这与普通的接收操作行为一致。

*/

var strChan = make(chan string, 3)

func main() {
	// 通道
	syncChan1 := make(chan struct{}, 1)
	// 通道
	syncChan2 := make(chan struct{}, 2)

	// goroutine#1
	go receive(strChan, syncChan1, syncChan2)
	// goroutine#2
	go send(strChan, syncChan1, syncChan2)

	<-syncChan2
	<-syncChan2
}

func receive(strChan <-chan string, syncChan1 <-chan struct{}, syncChan2 chan<- struct{}) {

	<-syncChan1
	fmt.Println("Received a sync signal and wait a second... [receiver]")
	time.Sleep(time.Second)
	for elem := range strChan {
		fmt.Println("Received:", elem, "[receiver]")
	}
	fmt.Println("Stopped. [receiver]")
	syncChan2 <- struct{}{}
}

func send(strChan chan<- string, syncChan1 chan<- struct{}, syncChan2 chan<- struct{}) {
	for _, elem := range []string{"a", "b", "c", "d"} {
		strChan <- elem
		fmt.Println("Sent:", elem, "[Sender]")
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
PS D:\gohome\houlin\ch4> go run .\chanbase3.go
Sent: a [Sender]
Sent: b [Sender]
Sent: c [Sender]
Sent a sync signal. [sender]
Received a sync signal and wait a second... [receiver]
Sent: d [Sender]
Wait 2 seconds... [sender]
Received: a [receiver]
Received: b [receiver]
Received: c [receiver]
Received: d [receiver]
Stopped. [receiver]
*/
