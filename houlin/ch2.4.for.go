package main

import "fmt"

/*

for 语句用于根据给定的条件，重复执行一个代码块。
这个条件或由 for 子句直接给出，或从 range 子句中获得。

*/

// for 子句
// 一条 for 语句可以携带一条 for 子句
// for 子句可以包含初始化子句、条件子句和后置子句

var number int

func generateNumber() {
	for i := 0; i < 100; i++ {
		fmt.Println(number)
		number++
	}
}

func for2() {
	var j uint = 1
	for ; j%5 != 0; j = j * 3 { // 省略初始化子句
		number++
		fmt.Println(number, j)
	}
}

func for3() {
	var m = 1
	for m < 50 { // 省略初始化子句和后置子句
		m *= 3
		fmt.Println(m)
	}
}

func for4() {
	ints := []int{1, 2, 3, 4, 5}
	for i, d := range ints {
		fmt.Printf("Index: %d, Value: %d\n", i, d)
	}
}
func main() {
	// generateNumber()
	// for2()
	// for3()
	for4()
}
