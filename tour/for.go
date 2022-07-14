package main

import "fmt"

func for_main() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Total sum:", sum)
}
