package main

import (
	"fmt"
	"sync"
	"time"
)

// 问题思考：为什么要在 main 函数执行结束的之前，使用了两次 time.Sleep()？
//
func main() {
	var mutex sync.Mutex
	fmt.Println("Lock the lock. (main)")
	mutex.Lock()
	fmt.Println("The lock is locked.(main)")
	// go func(){}() -> go 函数 启动 goroutine
	// 当主函数的例程运行结束了，基于主函数而产生的其他例程将被同样的结束掉。所以才需要用到 time.Sleep()
	for i := 1; i <= 3; i++ {
		go func(i int) {
			fmt.Printf("Lock the lock. (g%d)\n", i)
			mutex.Lock()
			fmt.Printf("The lock is locked. (g%d)\n", i)
		}(i)
	}
	time.Sleep(time.Second)                // main 函数睡眠，使得上面的 goroutine 有机会运行
	fmt.Println("Unlock the lock. (main)") // 睡醒之后，打印并解锁互斥锁
	mutex.Unlock()                         // 解锁互斥锁
	fmt.Println("The lock is unlocked. (main)")
	time.Sleep(time.Second) // 再次睡眠，等于其他 goroutine 运行
}

/*

PS D:\gohome\houlin\ch5> go run .\repeatedlylock.go
Lock the lock. (main)
The lock is locked.(main)
Lock the lock. (g1)
Lock the lock. (g3)
Lock the lock. (g2)
Unlock the lock. (main)
The lock is unlocked. (main)
The lock is locked. (g1)
*/
