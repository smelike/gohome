package main

import "fmt"

var intChan = make(chan int, 10)
var strChan = make(chan string, 10)

func main() {
	select {
	case e1 := <-intChan:
		fmt.Printf("The 1th case was selected. e1=%v.\n", e1)
	case e2 := <-strChan:
		fmt.Printf("The 2nd case was selected. e2=%v.\n", e2)
	default:
		fmt.Println("Default!")
	}
}

func sendInt(intChan chan<- int) {

}
