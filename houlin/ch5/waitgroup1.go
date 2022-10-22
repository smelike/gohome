package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	fmt.Println(time.Now())
	wg.Done()
	go func() {
		defer wg.Done() // 使用 defer 的效果等于在 func 结尾处调用 wg.Done()
		time.Sleep(time.Second * 2)
		// wg.Done()
	}()
	// wg.Done()
	// wg.Done()
	wg.Wait()
	fmt.Println(time.Now())
}
