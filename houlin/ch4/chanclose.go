package main

import "fmt"

// 关闭通道
/*
	向一个已关闭的通道发送元素值，会让发送操作引发运行时恐慌。

	关闭通道，要放松发送端。而不是接收端，接收端无法判断发送端是否还会向该通道发送元素值。
*/
func main() {
	dataChan := make(chan int, 5)
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)

	go func() { // 接收操作
		<-syncChan1 // 实现 goroutine 阻塞的目的
		for {
			if elem, ok := <-dataChan; ok {
				fmt.Printf("Received: %d [receiver]\n", elem)
			} else {
				break
			}
		}
		fmt.Println("Done. [receiver]")
		syncChan2 <- struct{}{}
	}()

	/*
		在向通道 dataChan 发送完所有元素值并关闭通道之后，才告知接收方开始接收。
		虽然通道关闭，但对于接收操作并无影响，接收方依然可以在接收完所有元素值后自行结束。

		注意：
		(1) 同一通道仅允许关闭一次，对通道的重复关闭会引发运行时恐慌；
		(2) 调用 close 函数时，需把代表[关闭的通道的变量]作为参数传入；如果变量的值为 nil，会引发运行时恐慌。
	*/
	go func() { // 发送操作
		for i := 0; i < 5; i++ {
			dataChan <- i
			fmt.Printf("Sent: %d [sender]\n", i)
		}
		close(dataChan)
		syncChan1 <- struct{}{}
		fmt.Println("Done. [sender]")
		syncChan2 <- struct{}{}
	}()
	<-syncChan2 // 目的：goroutine 阻塞操作的实现
	<-syncChan2

	// 长度与容量
	// 内建函数 len 和 cap 也是可以作用在通道上，作用是获取通道中当前的元素值数量（即长度）和通道可容纳元素值的最大数量（即容量）。
	// 通道的容量是在初始化时已经确定的，并且之后不能修改。
	// 而通道的长度则会随着实际情况变化。
	// 通过容量来判断是否带有缓冲，若其容量为 0，肯定就是一个非缓冲通道。容量大于 0 ，则就是一个缓冲通道。
}
