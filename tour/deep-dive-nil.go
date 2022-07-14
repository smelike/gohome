package main

import "fmt"

func test_main() {
	// nil := "this is nil"
	// fmt.Println(nil)
	var slice []string = make([]string, 0, 12)
	fmt.Println(slice)

	const val1 = iota
	fmt.Printf("%T\n", val1)
}
