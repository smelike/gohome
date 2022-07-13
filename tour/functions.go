package main

import (
	"fmt"
	"math/rand"
)

func d_add(x int, y int) int {
	return x + y
}

func d_main() {
	fmt.Println(add(100, 600))
	fmt.Println(add(rand.Intn(10), rand.Intn(100)))
}
