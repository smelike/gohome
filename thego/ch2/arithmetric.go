package main

import "fmt"

func main() {

	var u uint8 = 255
	fmt.Println(u, u+1, u*u) // 为什么 u*u 结果为 1？

	fmt.Println(16 &^ 8) // 0001 0000  0000 1000

	fmt.Println(1<<2, 1>>2)
}
