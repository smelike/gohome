package main

import "fmt"

func for_con_main() {
	sum := 1
	for sum < 1000 {
		fmt.Println("\n", sum)
		sum += sum
		// while will it ends
	}
	fmt.Println("Total sum: ", sum)
}
