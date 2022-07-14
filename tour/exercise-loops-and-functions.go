package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	n -= (n*n - x) / (2 * n)
	return n
}

var n = 1.0

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Pow(1.5, 2))
	fmt.Println(math.Sqrt(2))
}
