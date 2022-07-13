package main

import "fmt"

/*

two or more consecutive named function parameters share a type,
you can omit the type from all but the last.
In this example, we shortened

	x int, y int
to
	x, y int */

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("shortened: ", add(12, 43))
}
