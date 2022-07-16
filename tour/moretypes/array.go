package main

import "fmt"

/*
// Arrays

// The type [n]T is an array of n values of type T
// The expression: var a [10]int declare a variable a as an array of ten integers.
*/
func main() {

	// declare an array
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	fmt.Println("array a's length", len(a))

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	fmt.Println("array primes's length", len(primes))
}
