package main

import "os"

/*
单向通道(channel)

单向通道可分为发送通道和接收通道。

注意：无论哪一种单向通道，都不应出现在变量的声明中。
试想一下，如果声明并初始化了这样一个变量：var uselessChan chan<- int = make(chan<- int, 10)
那么应该怎样去使用它呢？显然，一个只进不出的通道没有任何意义。那单向通道的应用场景又在哪里？

问题：单向通道的应用场景在哪里？

单向通道应由双向通道变换而来。我们可以使用这种变换来约束程序对通道的使用方式。

[双向通道 -》变换约束 -》实现单向通道的切换]

os/signal.Notify 函数的声明：func Notify(c chan<- os.Signal, sig ...os.Signal)

该函数第一个参数的类型是发送通道类型。从表面上看，调用它的程序需要传入一个只能发送而不能接收的通道。
然而并不应该如此，在调用该函数时，你应该传入一个双向通道。Go 会依据该参数的声明，自动把它转换为单向通道。
Notify 函数中的代码只能向通道 c 发送元素值，而不能从它那里接收元素值。


*/

// 利用 Go 语言的语法规则做到了强约束。

// 接口声明：利用语法级别的约束避免实现类型对参数 c 进行错误的操作

// 单向通道版本 #1

type SignalNotifier1 interface {
	// c cha<- os.Signal 发送通道
	Notify(c chan<- os.Signal, sig ...os.Signal)
}

// 单向通道版本 #2
// 约束方法的调用方，Notify 方法的调用方只能从作为结果的通道中接收元素值，而不能向其发送元素值。
type SignalNotifier2 interface {
	Notify(sig ...os.Signal) <-chan os.Signal // 接收通道
}

// 方法的实现方 & 方法的调用方

/*
	版本#1 的方法声明更适合存在于接口类型中，因为它可以作为该接口的实现规则之一。
	type SignalNotifier interface {
		Notify(c chan<- os.Signal, sig ...os.Signal)
	}

	版本#2 的方法声明更适用于函数或结构体的方法，原因是它可以约束对函数或方法的结果值得使用方式。

	type SignalNotifier interface {
		Notify(sig ...os.Signal) <-chan os.Signal
	}

	该约束并不是绝对的。比如，在 os/Signal.Notify 函数的声明中，参数 c 的类型就隐含了函数调用方对该通道的使用规则。
	虽然此规则是可以轻易破坏的，但是这对于函数调用方来说没有任何好处。因此，这样是可以达到约束目的的。
	（规则是可以轻易破坏的，但这对于调用方来说没有任何好处）
*/
