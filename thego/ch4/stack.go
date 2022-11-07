package main

import "fmt"

/*
定义一个队列类型吗？
*/

func main() {

	s := []int{5, 6, 7, 8, 9}
	// fmt.Println(remove(s, 3))
	fmt.Println(remove2(s, 1))
}

// func stack(qs)

// remove the i-th element of a slice
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

// don't need to preserve the order, just move the last element into the gap.
func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1] // move the last element into the gap
	return slice[:len(slice)-1]
}
