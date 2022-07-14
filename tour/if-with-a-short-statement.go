package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func if_shortened_main() {
	fmt.Println(
		pow(3, 2, 10), // 3^2 < 10 return v
		pow(3, 3, 20), // 3^3 > 20 return lim
	)
}
