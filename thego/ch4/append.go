package main

import "fmt"

func main() {
	var runes []rune
	// []rune("Hello, 世界")
	for _, r := range "Drawing, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
}
