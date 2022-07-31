package main

import "fmt"

// 接口 Talk
type Talk interface {
	Hello(userName string) string
	Talk(heard string) (saying string, end bool, err error)
}

type myTalk string

func (talk *myTalk) Hello(userName string) string {

	h := fmt.Sprintf("Hello %s", userName)
	return h
}

func (talk *myTalk) Talk(heard string) (saying string, end bool, err error) {
	saying = fmt.Sprintf("Talk: %s", heard)
	return
}

//备注：myTalk 类型并不是 Talk 接口的实现类型， *myTalk 类型才是 Talk 接口的实现类型

// 接口 Chatbot
type Chatbot interface {
	Name() string
	Begin() (string, error)
	Talk
	ReportError(err error) string
	End() error
}

// 结构体 simpleCN
type simpleCN struct {
	name string
	talk Talk
}
