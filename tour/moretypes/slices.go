package main

import "fmt"

// Slices

// An array has a fixed size.
// A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.
// In practice, slices are much more common than arrays.

// The type []T is a slice with elements of type T.

// A slice is formed by specifying two indices, a low and high bound, separated by a colon: a[low: high]

// http://127.0.0.1:3999/moretypes/7

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13} // array

	var s []int = primes[1:4] // slice

	fmt.Println(s)
}
