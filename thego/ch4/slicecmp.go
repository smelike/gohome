package main

import "fmt"

func main() {

	x := []string{"aa", "bb", "ccc"}
	y := []string{"aa", "bb", "cccd"}

	fmt.Printf("%t", equal(x, y)) // %t to show boolean
}

/*
比较两个字符串切片
*/
func equal(x, y []string) bool {
	// 长度比较
	if len(x) != len(y) {
		return false
	}
	// 字符比较
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
