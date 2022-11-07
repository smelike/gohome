package main

import "fmt"

func main() {

	// ages := make(map[string]int) // mapping from strings to ints

	// a map literal
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}

	// equal to
	// ages := make(map[string]int)
	// ages["alice"] = 31
	// ages["charlie"] = 34

	// empty map map[string]int{}
	fmt.Println(ages)
	fmt.Println(ages["alice"])
	delete(ages, "alice")

	ages["bob"] = ages["bob"] + 1
	fmt.Println(ages["alice"], ages["bob"], ages["dell"])

	x := []int{}
	z = x[:1]
	fmt.Println(z)
}
