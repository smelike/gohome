package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1) // byte() convertion i 位与 1 的结果
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCount2(x uint64) int {
	var t int = 0
	for i := 0; i < 8; i++ {
		t += int(pc[byte(x>>(i*8))])
	}
	return t
}

/*
Note that the range loop in init uses only the index; the value is unnecessary and thus
need not be included.
The loop could also have been written as `for i, _ :=range pc{}`
*/

/*
Exercise 2.3 Rewrite PopCount to use a loop instead of a single expression. Compare
the performance of the two version.

Exercise 2.4 Write a version of PopCount that counts bits by shifting its argument through 64
bit position, testing the rightmost bit each time. Compare its performance to the table-lookup version.

Exercise 2.5 The expression x&(x-1) clears the rightmost non-zero bit of x. Write a version of PopCount
that counts bits by using this fact, and assess its performance.

*/
