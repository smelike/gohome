package main

import (
	"errors"
	"fmt"
)

func main() {

	var ipv4 [4]uint8 = [4]uint8{192, 168, 0, 1}
	fmt.Println(ipv4)

	var ips = []string{"192.168.0.1.", "192.168.0.2", "192.168.0.3"}
	fmt.Println("Slice at initialize:", ips)

	// ips[:cap(ips)]
	ips = append(ips, "192.18.0.4")

	fmt.Println("slice after append:", ips)

	// 实验：引用 operate，返回结果是 nil
	/* result, err := operate(20, 10, divide)
	if err != nil {
		fmt.Println("divide result is:", result)
	}
	fmt.Println(err) */
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

// 定义二元操作的函数类型
type binaryOperation func(operand1 int, operand2 int) (result int, err error)

// 进一步范化：用于以自定义的方式执行二元操作
func operate(op1 int, op2 int, bop binaryOperation) (result int, err error) {
	if bop == nil {
		err = errors.New("Invalid binary operation function")
		return
	}
	return bop(op1, op2)
}
