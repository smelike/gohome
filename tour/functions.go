package main

import (
	"fmt"
	"math/rand"
)

func d_add(x int, y int) int {
	return x + y
}

func d_main() {
	fmt.Println(d_add(100, 600))
	fmt.Println(d_add(rand.Intn(10), rand.Intn(100)))
}
