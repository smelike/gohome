package main

import (
	"fmt"
)

/*

（1）接口：用于定义一组行为
每个行为都由一个方法声明表示。
接口类型中的方法声明只有方法签名而没有方法体，而方法签名包括且包括方法的名称、参数列表和结果列表。
举个例子，如果要定义“聊天”相关的一组行为，可以这样写：

*/

type Talk interface {
	Hello(userName string) string
	Talk(heard string) (saying string, end bool, err error)
}

/*

type、接口类型名称、interface 以及由花括号包裹的方法声明集合，共同组成了一个接口类型声明。
注意，其中每个方法声明必须独占一行。

*/

/*
只要一个数据类型的方法集合中包含 Talk 接口声明的所有方法，那么它就一定是 Talk 接口的实现类型。
显然，这种接口实现方式完全是非侵入式的。

Talk 接口的实现类型可以是这样的：
*/

type myTalk string

func (talk *myTalk) Hello(userName string) string {
	// code to be executed
	fmt.Println(userName)
	err := "something error"
	return err
}

func (talk *myTalk) Talk(heard string) (saying string, end bool, err error) {
	fmt.Println("myTalk function Talk()")
	// saying := "hello"
	// end := true
	// err := errors.New("function is error")
	return
}

// 实例化
var talk Talk = new(myTalk)

/*

Go 的数据类型之间并不存在继承关系，接口类型之间也是如此。
一个接口类型的声明中可以嵌入任意其他接口类型。
一组行为中可以包含其他的行为组，而且数量不限。
下面 Chatbot 的声明中就嵌入了 Talk 接口类型：
*/
type Chatbot interface {
	Name() string
	Begin() (string, error)
	Talk
	ReportError(err error) string
	End() error
}

func main() {
	fmt.Println(talk)
	// fmt.Printf("%v", talk)
	ret := talk.Hello("jaemsxu")
	fmt.Println(ret)
}
