package main

import (
	"fmt"
	"time"
)

func main() {

	// <初始化的绝对时间> + <相对到期时间> == <绝对到期时间>
	timer1 := time.NewTimer(time.Second)

	fmt.Printf("Present time: %v. \n", time.Now())

	expirationTime := <-timer1.C
	fmt.Printf("Expiration time: %v. \n", expirationTime)
	fmt.Printf("Stop timer: %v. \n", timer1.Stop())
	// 3 小时 36 分钟的定时器
	// timer2 := time.NewTimer(3*time.Hour + 36*time.Minute)

}
