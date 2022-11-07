package main

import "fmt"

func main() {

	ages := make(map[string]int)
	fmt.Println(ages)

	// subscripting a map
	if _, ok := ages["foo"]; !ok {
		fmt.Println("foo key is not exist..")
	}

	collect1 := map[string]float32{
		"A book": 3.4,
		"B book": 5.0,
		"C book": 4.3,
	}
	collect2 := map[string]float32{
		"A book": 3.4,
		"B book": 5.0,
		"C book": 4.3,
	}

	fmt.Println(equal(collect1, collect2))

	fmt.Println(
		equal(map[string]float32{"A": 42},
			map[string]float32{"A": 42}))
}

func equal(x, y map[string]float32) bool {
	if len(x) != len(y) {
		return false
	}

	// natively xv != y[k], omit the key-exist?
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
