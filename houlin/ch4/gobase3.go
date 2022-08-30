package main

import (
	"fmt"
	"time"
)

func main() {
	name := "Eric"
	go func() {
		fmt.Printf("Hello, %s!\n", name)
	}()
	time.Sleep(time.Millisecond)
	name = "Harry"
	// runtime.Gosched()
}
