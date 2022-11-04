package main

import "fmt"

func main() {
	// []int slice no fixed size
	// var s []int = []int{12, 34, 56, 78, 90} // slice
	s := [...]int{0, 1, 2, 3, 4, 5} // array
	n, err := fmt.Println(rev(s[:]), len(s))
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("bytes by written: %d", n)
}

func rev(s []int) []int {

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
