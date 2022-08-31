package main

import (
	"fmt"
	"time"
)

// 要同时问候多个人，名单如下：names := []string{"Eric", "Robert", "Jim", "Mark"}

func main() {
	names := []string{"Eric", "Harry", "Robert", "Jim", "Mark"}

	for _, name := range names {
		go func() {
			fmt.Printf("Hello, %s!\n", name)
		}()
		time.Sleep(time.Millisecond)
	}
	// time.Sleep(time.Millisecond)
	// go 函数的执行时机，Go 的调度器

}
