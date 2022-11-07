package main

import "fmt"

func main() {
	var intValus []int = []int{12, 34, 56, 78, 90}
	iterate(intValus...)
}

func iterate(y ...int) {
	for _, v := range y {
		fmt.Printf("value = %v\n", v)
	}
}
