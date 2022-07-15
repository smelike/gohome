package main

import (
	"fmt"
	"time"
)

// Switch with no condition
// Switch without a condition is the same as `switch true`
// This construct can be a clean way to write long if-then-else chains.
func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good Evening!")
	}
}
