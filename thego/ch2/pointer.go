package main

import "fmt"

var p = f()

func main() {
	fmt.Println(p)
	fmt.Println(f() == f())
}

/*
	func f() *int { // pointer to an integer
		v := 1
		return &v
	}
*/
func f() *string { // pointer to an string
	v := "string"
	return &v
}
