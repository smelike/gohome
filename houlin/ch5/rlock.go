package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var rwm sync.RWMutex
	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Printf("Try to lock for reading... [%d]\n", i)
			rwm.RLock()
			fmt.Printf("Locked for reading. [%d]\n", i)
			time.Sleep(time.Second * 2)
			fmt.Printf("Try to unlock for reading... [%d]\n", i)
			rwm.RUnlock()
			fmt.Printf("Unlocked for readings. [%d]\n", i)
		}(i)
	}
	time.Sleep(time.Millisecond * 100) // 睡眠等待运行时系统运行 goroutine
	fmt.Println("Try to lock for writing...")
	fmt.Println(time.Now())
	rwm.Lock()
	fmt.Println(time.Now())
	fmt.Println("Locked for writing.")
}

/*
go run .\rlock.go
Try to lock for reading... [0]
Locked for reading. [0]
Try to lock for reading... [1]
Locked for reading. [1]
Try to lock for reading... [2]
Locked for reading. [2]
Try to lock for writing...
Try to unlock for reading... [2]
Unlocked for readings. [2]
Try to unlock for reading... [1]
Unlocked for readings. [1]
Try to unlock for reading... [0]
Unlocked for readings. [0]
Locked for writing.

// 读操作的锁定操作，使得主 goroutine 阻塞，因为读锁定未解锁，经过 time.Sleep(2s)  2s 后

*/

/*
Try to lock for reading... [0]
Locked for reading. [0]
Try to lock for reading... [2]
Try to lock for reading... [1]
Locked for reading. [1]
Locked for reading. [2]
Try to lock for writing...
2022-10-09 12:59:57.6116138 +0800 CST m=+0.109863001 [写锁定被阻塞的时间]
Try to unlock for reading... [2]
Unlocked for readings. [2]
Try to unlock for reading... [1]
Unlocked for readings. [1]
Try to unlock for reading... [0]
Unlocked for readings. [0]
2022-10-09 12:59:59.5167679 +0800 CST m=+2.015013101 [写锁定操作成功的时间]
Locked for writing.
*/
