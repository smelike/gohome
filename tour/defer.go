package main

import (
	"fmt"
)

// A defer statement defers the execution of a function until the surrounding function returns.

func defer_main() {

	// block one
	// deferBlock1()

	//block two
	// deferBlock2()

	var name string = "old value"
	defer fmt.Println(name) // be executed after Line 21, even variable name's value has been changed. But the print still is the previous value

	name = "change to new value"
	fmt.Println(name)
}

func deferBlock1() {
	fmt.Println("----deferBlock#1------")
	fmt.Printf("\n")
	defer fmt.Println("world")
	fmt.Println("hello")
	fmt.Println("------end-deferBlock#1--------")
	fmt.Printf("\n")
}
func deferBlock2() {
	fmt.Println("--------deferBlock#2---------")
	fmt.Printf("\n")
	var name string = "old value"
	defer fmt.Println(name)

	name = "change to new value"
	fmt.Println(name)
	fmt.Println("---------end#2---------------")

}
