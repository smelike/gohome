package main

import "fmt"

func main() {
	var i int
	j := i // j is an int

	// := new assignment, right side must be variable
	i = 42            // int
	f := 3.142        // float
	g := 0.867 + 0.5i // complex128

	fmt.Printf("%T, %T, %T, %T", i, j, f, g)
}
