package main

import "fmt"

// 表达式是把操作符和函数作用于操作数的计算方法

// 方法是函数的一种，实际是与某个数据类型关联在一起的函数

type myInt int

// 值方法的接收者类型

/* func (i myInt) add(another int) myInt {
	i = i + myInt(another)
	return i
} */

// 指针方法的接收者类型
func (i *myInt) add(another int) myInt {
	*i = *i + myInt(another)
	// *i 表示取值操作
	// &i 表示取址操作
	return *i
}

/*
从声明上，方法只是在关键字 func 和函数名称之间，加了一个由圆括号包裹的接收者声明。——接收者声明

接收者声明，由两部分组成：
- 右边表明这个方法与哪个类型关联，这里是 myInt；
- 左边指定这个类型的值在当前方法中的标识符，这里是 i。
这个标识符在当前方法中可以看作一个变量的代表，就像参数那样。
所以，它可以称为接收者变量。

*/

func main() {
	i1 := myInt(1)  // i1 源值
	i2 := i1.add(2) // i1.add()方法内的 i 是接收者变量，i 获得了 i1 的副本
	fmt.Println(i1, i2)
	/* 这三行代码执行后，会打印出 1 3。
	i的值未改变，是因为在值方法中对接收者变量的赋值一般不会影响到源值。
	这里，变量 i1 的值就是源值。在调用 i1 的 add 方法时，这个值被赋给了接收者变量 i(前者的副本与后者产生关联)。
	但是，i 和 i1 是两个变量，它们之间并不存在关联。 */
}
