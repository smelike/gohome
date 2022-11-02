package main

import "fmt"

func main() {
	fmt.Printf("case %d, return %d\n", -10, signum(-10))
	fmt.Printf("case %d, return %d\n", 10, signum(10))
	fmt.Printf("case %d, return %d\n", 0, signum(0))
}

// compare number
func signum(x int) int {
	switch { // equivalent to switch true
	case x > 0:
		return +1
	default:
		return 0
	case x < 0:
		return -1
	}
}
