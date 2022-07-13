package main

import "fmt"

/*

two or more consecutive named function parameters share a type,
you can omit the type from all but the last.
In this example, we shortened

	x int, y int
to
	x, y int */

func e_add(x, y int) int {
	return x + y
}

func e_main() {
	fmt.Println("shortened: ", e_add(12, 43))
}
