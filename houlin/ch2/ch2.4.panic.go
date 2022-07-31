package main

import (
	"errors"
	"fmt"
)

func main() {
	// outerFunc()

	myIndex := 4
	ia := [3]int{1, 2, 3}
	_ = ia[myIndex]
	/*
		panic: runtime error: index out of range [4] with length 3

		goroutine 1 [running]:
		main.main()
				D:/gohome/houlin/ch2.4.panic.go:9 +0x1d
		exit status 2 */
}

func outerFunc() {
	innerFunc()
}

func innerFunc() {
	panic(errors.New("An intended fatal error!"))
}

func printNumbers() {
	for i := 0; i < 5; i++ {
		defer func(n int) {
			if p := recover(); p != nil {
				fmt.Printf("Recovered panic: %s\n", p)
			}
			fmt.Printf("%d", n)
		}(i)
	}
}
