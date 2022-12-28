package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	t, z := 0., 1.
	for {
		z, t = z-(z*z-x)/(2*z), z
		if math.Abs(t-z) < 1e-8 {
			break
		}
	}
	return z
}

func main() {
	i := 169.
	fmt.Println(Sqrt(i))
	fmt.Println(Sqrt(i) == math.Sqrt(i))
}
