package main

import "fmt"

// Variadic Functions

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func max(vals ...int) int {
	if len(vals) == 0 {
		fmt.Println("Parameters cannot be empty.")
		//return nil
	}
	max := 0
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func min(vals ...int) int {
	min := 0
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}
func main() {
	fmt.Println(sum())
	fmt.Println(sum(3))
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum([]int{10, 20, 30, 40}...))
	fmt.Println(max([]int{10, 30, 200}...))
}
