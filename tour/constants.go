package main

import "fmt"

/*

	Constants are declared like variables, but with the const keyword.

	Constants can be character, string, boolean, or numeric values.

	Constants cannot be declared using the := syntax.

*/
const Pi = 3.14

func o_main() {
	const world = "世界"
	fmt.Println("Hello ", world)
	fmt.Println("Happy ", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
