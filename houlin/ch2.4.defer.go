package main

import "fmt"

func outerFunc() {
	defer fmt.Println("函数执行结束前一刻才会被打印。")
	fmt.Println("第一个被打印")
}

// without pass params from outside
func printNumbers() {
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Printf("%d", i)
			// 55555，由延迟函数的执行时机引起的。
			// 待那 5 个延迟函数执行时，它们使用的 i 值已经是 5 了。
		}()
	}
}

func printNumbers2() {
	for i := 0; i < 5; i++ {
		defer func(n int) {
			fmt.Printf("%d", n)
		}(i)
	}
}
func main() {
	outerFunc()
	fmt.Println("---delemeter-----")
	printNumbers()
	fmt.Println("---delemeter-----")
	printNumbers2()
}
