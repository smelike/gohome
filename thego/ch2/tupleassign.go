package main

import "fmt"

// func tempF() interface{}
// type tempF interface {}
func main() {
	fmt.Println(gcd(10, 2))
	fmt.Println(fib(3))

	// var tempF int = 212
	// res, ok = tempF.(int)
	// fmt.Println(res, ok)
}

// greates common divisor
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
		// fmt.Println(x, "=", y)
	}
	return y
}
