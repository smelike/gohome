package main

/*
并发安全的整数数组类型，无锁化方案使用原子值实现
*/
type ConcurrentArray interface {
	Set(index uint32, elem int) (err error)
	Get(index uint32) (elem int, err error)
	Len() uint32
}
