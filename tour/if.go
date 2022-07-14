package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {

	// if x's value small 0
	if x < 0 {
		// 自调用 sqrt(-x)
		return sqrt(-x) + "i"
		// i 是虚数后序
	}

	// math 包的方法 math.Sqrt()
	return fmt.Sprint(math.Sqrt(x))
}

func if_main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
