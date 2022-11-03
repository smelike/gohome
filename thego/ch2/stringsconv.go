package main

import "fmt"

func main() {
	// s := "abc"
	s := "123456789"
	b := []byte(s)
	s2 := string(b)

	fmt.Println(s, b, s2)
}
