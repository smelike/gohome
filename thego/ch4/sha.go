package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var sha = flag.Int("sha", 256, "hash length, such as sha256...")

func main() {
	flag.Parse()
	var s [64]byte
	var x string = "x"
	switch *sha {
	default:
		s = sha256.Sum256([]byte(x)) // sha256.Sum256() 返回 [32]byte
	case 512:
		s = sha512.Sum512([]byte(x)) // 返回 [64]byte
	}

	fmt.Println(s)
}
