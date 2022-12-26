package main

import (
	"fmt"
)

const Delta = 0.0001

func isConverged(d float64) bool {
	if d < 0.0 {
		d = -d
	}
	if d < Delta {
		return true
	}
	return false
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	tmp := 0.0
	for {
		tmp = z - (z*z-x)/2*z
		if d := tmp - z; isConverged(d) {
			return tmp, nil
		}
		z = tmp
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	//return fmt.Sprint(e)
	/*
			runtime: goroutine stack exceeds 1000000000-byte limit
		runtime: sp=0xc020161380 stack=[0xc020160000, 0xc040160000]
		fatal error: stack overflow
	*/
	return fmt.Sprintf("cannot Sqrt negative number: %f\n", e)
}

func main() {
	fmt.Println(Sqrt(2))
	if _, err := Sqrt(-2); err != nil {
		fmt.Println(err)
	}
}
