package main

import "fmt"

// rotate a slice left by n elements
func main() {

	s := rotaten([]int{1, 2, 3, 4, 5, 6}, 3)
	fmt.Println(s)
}

func rotaten(s []int, n int) []int {
	if n > 0 && n <= len(s) {
		rev(s[:n])
		rev(s[n:])
		return rev(s)
	}
	return nil
}
func rev(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
