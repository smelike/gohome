package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	e19()
	if r := recover(); r != nil {
		fmt.Println(r)
	}
	const day = 24 * time.Hour
	fmt.Println(day.Seconds())
}

func e19() {
	defer func() {
		e := recover()
		// fmt.Println(e)
		if e != nil {
			panic(e)
		}
	}()
	panic(math.MaxInt)
}
