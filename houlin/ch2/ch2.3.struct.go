package main

// struct - 结构体
/*

结构体类型不仅可以关联方法，而且可以有内置元素（又称字段）。
结构体类型的声明一般以关键字 type 开始， 并依次包含类型名称、关键字 struct 以及由花括号包裹的字段声明列表。

*/

// 接口
type Talk interface {
	Hello(userName string) string
	Talk(heard string) (saying string, end bool, err error)
}

// 结构体
type simpleCN struct {
	name string
	talk Talk
}
