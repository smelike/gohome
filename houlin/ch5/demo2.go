package cow

import (
	"errors"
	"fmt"
	"sync/atomic"
)

/*
并发安全的整数数组类型，无锁化方案使用原子值实现
*/

// 先定义接口
/*
	整型数组的长度是固定的，并且必须在初始化时给定。
	对该数组的变更仅限于对其中某个元素值的修改。
*/
type ConcurrentArray interface {
	Set(index uint32, elem int) (err error)
	Get(index uint32) (elem int, err error)
	Len() uint32
}

// 该接口的实现类型的基本声明
// concurrentArray 用于表示 ConcurrentArray 接口的实现类型
/*
这个结构体类型的字段只有两个，分别代表数组长度和支持原子操作的值。
*/
type concurrentArray struct {
	length uint32
	val    atomic.Value
}

// 用于创建整型数组值的 NewConcurrentArray 函数
// 创建一个 ConcurrentArray 类型值
func NewConcurrentArray(length uint32) ConcurrentArray {
	array := concurrentArray{}
	array.length = length
	// 存储在字段 val 中的是一个切片值，而不是数组值。
	// mkae([]int, array.length) 创建的是切片，还是数组？
	array.val.Store(make([]int, array.length))
	return &array
}

// 指针方法的 Set
func (array *concurrentArray) Set(index uint32, elem int) (err error) {

	// 检查索引
	if err = array.checkIndex(index); err != nil {
		return
	}

	// 检查舒数组的元素值
	if err = array.checkValue(); err != nil {
		return
	}

	newArray := make([]int, array.length)
	// array.val.Load().([]int) // 将 Load() 返回值格式化为 []int
	// Load() returns the value set by the most recent Store
	copy(newArray, array.val.Load().([]int))
	newArray[index] = elem
	array.val.Store(newArray)
	return
}

func (array *concurrentArray) Get(index uint32) (elem int, err error) {

	// 检查是否存在有 index，有即返回数组中的元素
	if err = array.checkIndex(index); err != nil {
		return
	}
	if err = array.checkValue(); err != nil {
		return
	}
	elem = array.val.Load().([]int)[index]
	return
}

func (array *concurrentArray) Len() (len uint32) {
	return array.length
}

// 检查索引是否超出数组的长度
func (array *concurrentArray) checkIndex(index uint32) (err error) {
	if index >= array.length {
		return fmt.Errorf("Index out of range [0, %d]!", array.length)
	}
	return nil
}

func (array *concurrentArray) checkValue() (err error) {
	v := array.val.Load()
	if v == nil {
		return errors.New("Invalid int array!")
	}
	return nil
}
