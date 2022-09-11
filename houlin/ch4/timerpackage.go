package main

/*
	time 包

	time 包中的定时器（time.Timer） 和断续器（time.Ticker）都充分利用了缓冲通道的异步特性来传送到期通知。
	（缓冲通道的异步特性、传送到期通知）

	定时器 timer 应用的场景：
		利用定时器设定某一个任务的超时时间，这相当于对它们完成时间点进行控制。
		（任务的超时时间，等于控制任务的完成时间点。）

	断续器 ticker 应用的场景：
		常用断续器来设定任务的开始时间点。

	从这个角度看，它们面向的是两个看似对立又相互关联的方面。通过对它们的组合使用，可以有效控制对时间敏感的流程。

	ticker -> timer
	(开始时间 -> 超时时间[结束时间])

*/
