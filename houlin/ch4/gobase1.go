package main

import "runtime"

func main() {
	go println("Go!.Goroutine!")
	/*

		go println() 这行内容并不会出现。这是为什么呢？
		运行时系统会并发地执行 go 函数。运行时系统会使用一个 G 封装 go 函数并把它放到可运行 G 队列中，但是至于这个新的 G 什么时候会运行，就要看调度器的实时调度情况了。
		当该 go 语句之后没有任何语句。一旦 main 函数执行结束，就意味着该 Go 程序运行的结束。可是，这个时候那个新的 G 还没来得及执行。
		这种情况几乎总是会发生，所以我们不要对这种并发执行的先后顺序有任何假设，也不要指望 main 函数所在的 G 会最后一个运行完毕。
	*/

	// 干预多个 G 的执行顺序，最简陋的一种方法是使用 time 包中的 Sleep 函数
	// time.Sleep 的作用是让调用它的 goroutine 暂停（进入 Gwaiting 状态）一段时间。
	// time.Sleep(time.Millisecond)
	// runtime.Gosched() 暂停当前的 G
	runtime.Gosched()
}
