package main

import (
	"fmt"
	"os"
)

func main() {
	// var s, sep string
	// 循环、字符拼接
	// start := time.Now()
	for index, arg := range os.Args[1:] {
		// s += sep + arg
		// sep = " "
		fmt.Printf("%d:%s\n", index, arg)
		// fmt.Sprintf()
	}
	// fmt.Println(s)
}
