package main

import (
	"fmt"
	"math"
)

// the if statement can start with a short statement to execute before the condition.
// variables declared by the statement are only in scope until the end of the if.
func if_with_pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
	// return v // undefined: v
}

func if_with_a_short_statement_main() {
	fmt.Println(
		pow(3, 2, 10), // 3^2 < 10 return v
		pow(3, 3, 20), // 3^3 > 20 return lim
	)
}
