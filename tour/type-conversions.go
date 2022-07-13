package main

import (
	"fmt"
	"math"
)

// the expression T(v) converts the value v to the type T
func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*y + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
