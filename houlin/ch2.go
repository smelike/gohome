package main

import (
	"errors"
	"fmt"
)

func main() {

	var ipv4 [4]uint8 = [4]uint8{192, 168, 0, 1}
	fmt.Println(ipv4)

	var ips = []string{"192.168.0.1.", "192.168.0.2", "192.168.0.3"}
	fmt.Println(ips)

	// ips[:cap(ips)]
	ips = append(ips, "192.18.0.4")
}

// 如果函数声明的结果是有名称的，那么 return 关键字后面就不用追加任何东西了。
// Go 编程惯用法，即把 error 类型的结果作为函数结果列表的最后一员。
func divide(dividend int, divisor int) (result int, err error) {
	if divisor == 0 {
		err = errors.New("division by zero")
		return
	}
	result = dividend / divisor
	return
}
