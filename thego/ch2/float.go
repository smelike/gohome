package main

import (
	"fmt"
	"math"
)

// %8.3f aligned in an eight-character field
func main() {
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d ex = %8.3f\n", x, math.Exp(float64(x)))
	}

	x := math.Sqrt(-1)
	fmt.Println(math.NaN(), math.IsNaN(x))
}
