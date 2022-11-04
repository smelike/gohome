package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	fmt.Println(hashdiff("x", "X"))
}

func hashdiff(v1, v2 string) int {

	c1 := sha256.Sum256([]byte(v1))
	c2 := sha256.Sum256([]byte(v2))

	fmt.Printf("%x\n%x", c1, c2)
	var n = 0
	for i, v := range c1 {
		if v != c2[i] {
			n++
		}
	}
	return n
}
