package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s := "Hello, 世界" // 9 runes

	fmt.Println(len(s)) // 13 bytes
	fmt.Println(utf8.RuneCountInString(s))

	for i := 0; i < len(s); {
		// size, the number of bytes occupied by the UTF-8 encoding of r
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c - %d\n", i, r, size)
		i += size
	}

	var rl = utf8.RuneCountInString(s)
	for i := 0; i < rl; i++ {
		fmt.Println(s[i])
	}

	for i, r := range s {
		fmt.Printf("%d\t%[2]q\t%[2]d\n", i, r)
	}

	n := 0

	/* for _, _ = range s {
		n++
	} */
	for range s {
		n++
	}
	println(n)
}
