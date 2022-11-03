package main

import (
	"fmt"
)

/*
	每 3 位数字前加逗号
*/
func main() {
	var s string = "12345678"

	fmt.Println(comma(s))
}

// 递归调用函数
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	// return 需要不断处理的字符串
	return comma(s[:n-3]) + "," + s[n-3:]
}
