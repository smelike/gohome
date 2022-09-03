package main

import "fmt"

/*
	单向通道通常由双向通道转换而来，那么，单向通道是否可以转换回双向通道呢？
	答案是否定的。

	通道允许的数据传递方向是其类型的以部分。

	对于两个通道类型而言，数据传递方向的不同就意味着它们类型的不同。


*/

func main() {
	var ok bool

	// 双向通道
	ch := make(chan int, 1)

	_, ok = interface{}(ch).(<-chan int)
	fmt.Println("chan int => <-chan int:", ok)

	//发送通道
	sch := make(chan<- int, 1)
	_, ok = interface{}(sch).(chan int)
	fmt.Println("chan<- int => chan int:", ok)

	// 接收通道
	rch := make(<-chan int, 1)
	_, ok = interface{}(rch).(chan int)
	fmt.Println("<-chan int => chan int:", ok)
}
