package main

import "fmt"

func main() {
	var a [3]int             // array of 3 integers
	fmt.Println(a[0])        // print the first element
	fmt.Println(a[len(a)-1]) // print the last element

	// Print the indices and elements

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// print the elements only
	for _, v := range a {
		fmt.Printf("%d \n", v)
	}

	// an array literal to initialize an array with a list of values:
	var q [3]int = [3]int{22, 44, 66}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2])
	fmt.Println(q)

	// use "..."

	q = [...]int{1, 2, 3}
	fmt.Println(q)

	// specify a list of index and value pairs, like this:

	type Currency int

	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(RMB, symbol[GBP], symbol[RMB])

	ints := [...]int{99: -1}
	fmt.Println(len(r), ints[99], ints[98])
}
