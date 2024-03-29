package main

import "fmt"

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i) // 往 x []int 插入新元素，并返回新的 slice
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y // 更新后的 y 赋值给 x，才能进行下一次的 append 操作
	}

	/* var runes []int = []int{12, 34, 56, 78, 90}
	runes = append(runes, 2000)
	fmt.Printf("%v\t length=%d\t cap=%d\n", runes, len(runes), cap(runes))
	runes = append(runes, runes...) // append the slice runes
	fmt.Println(runes)
	fmt.Println(runes) */
}

func appendInt(x []int, y int) []int {
	var z []int

	/*
		// 当 x []int has zero values，必须加 1，
		zlen = len(x) -> panic: runtime error:
		index out of range [0] with length 0
	*/
	zlen := len(x) + 1 // why add 1?? because x's initializer is empty slice, has zero values.
	// zlen := len(x)
	if zlen <= cap(x) {
		// there is room is to grow. Extend the slice.
		z = x[:zlen] // z and x have the same underlying array
		fmt.Println("zlen<=cap:", z)
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x) // twice expand capacity
		}
		// 调整 capacity
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y // 添加到 len(x) index(index 的值范围其实是 0 - len(x) - 1) 处？
	return z
}

func appendInt2(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {

	} else {

	}
	copy(z[len(x):], y)
	return z
}
