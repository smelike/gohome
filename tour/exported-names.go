package main

import (
	"fmt"
	"math"
)

// exported names
func c_main() {
	/*
		a name is exported if it begins with a capital letter
		For example, Pizza is an exported name, as is Pi, which is exported from the math package
	*/
	// fmt.Println(math.pi)
	fmt.Println("the Pi number in math is ", math.Pi)
}
