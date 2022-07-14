package main

import "fmt"

func forwhile_main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("for is gos while, total sum:", sum)
}
