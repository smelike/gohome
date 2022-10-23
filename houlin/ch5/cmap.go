package cmap

import (
	"math"
	"sync/atomic"
)

/*
	散列段
	分段锁
	字典分段
	分段设计
	段定位算法：根据键来决定该键-元素对应该放入字典中的哪个散列段的计算方法。

	若不在并发安全字典中分段会怎样？
*/

// 先要确定并发安全的字典类型的行为
// 确定行为，显然需要一个接口类型

// 行为被接口类型定义
type ConcurrentMap interface {
	// 返回并发量
	Concurreny() int
	//  键-元素对
	Put(key string, element interface{}) (bool, error)
	// 获取与指定键关联的那个元素
	Get(key string) interface{}
	// 删除键
	Delete(key string) bool
	// 返回当前指定中键-元素对的数量
	Len() uint64
}

// 空接口类型的元素 是什么意思
// string 类型的键

// 接口类型的实现类型，这里使用结构体类型

type myConcurrentMap struct {
	concurrency int
	segments    []Segment // type Segment interface {...}
	total       uint64
}

// 实例函数：创建 ConcurrentMap 类型的实例
/*
PairRedistributor 代表一个接口类型
在散列段中做键-元素对的负载均衡

MAX_CONCURRENDY 允许的最大并发量
DEFAULT_BUCKET_NUMBER 一个散列段中默认包含的散列的数量

myConcurrentMap 的指针类型是 ConcurrentMap 接口的实现类型
*/
func NewConcurrentMap(
	concurrency int,
	pairRedistributor PairRedistributor) (ConcurrentMap, error) {
	if concurrency <= 0 {
		return nil, newIllegalParameterError("concurrency is too small")
	}
	if concurrency > MAX_CONCURRENCY {
		return nil, newIllegalParameterError("concurrency is too large")
	}
	cmap := &myConcurrentMap{}
	cmap.concurrency = concurrency
	cmap.segments = make([]Segment, concurrency)
	for i := 0; i < concurrency; i++ {
		cmap.segments[i] = newSegment(DEFAULT_BUCKET_NUMBER, pairRedistributor)
	}
	return cmap, nil
}

func (cmap *myConcurrentMap) Concurrency() int {
	return cmap.concurrency
}

/*
Pair 类型实际上是一个接口
Pair 接口首先嵌入了 linkedPair 接口，后者（linkedPair）是包级私有的，主要是为了保护
一些需要接口化的方法，使之不被包外代码访问。实现 linkedPair 接口，可以让多个键-元素对
形成一个单链表。
【单链表】
之所以有 Hash 方法，原因是：一个键-元素对值得键不可改变。因此，其键得散列值也是永不变得。
因此，在创建键-元素对值的时候，先计算出这个散列值并存储起来以备后用。这样可以节省一些后续计算，提高效率。


*/
func (cmap *myConcurrentMap) Put(key string, element interface{}) (bool, error) {
	p, err := newPair(key, element)
	if err != nil {
		return false, err
	}
	s := cmap.findSegment(p.Hash())
	ok, err := s.Put(p)
	if ok {
		atomic.AddUint64(&cmap.total, 1)
	}
	return ok, err
}

func (cmap *myConcurrentMap) Get(key string) interface{} {
	keyHash := hash(key)
	s := cmap.findSegment(keyHash)
	pair := s.GetWithHash(key, keyHash)
	if pair == nil {
		return nil
	}
	return pair.Element()
}
func (cmap *myConcurrentMap) Delete(key string) bool {
	s := cmap.findSegment(hash(key))
	if s.Delete(key) {
		atomic.AddUint64(&cmap.total, ^uint64(0))
		return true
	}
	return false
}

func (cmap *myConcurrentMap) Len() uint64 {
	return atomic.LoadUint64(&cmap.total)
}

// findSegment 会根据给定参数寻找并返回对应散列段
func (cmap *myConcurrentMap) findSegment(keyHash uint64) Segment {
	if cmap.concurrency == 1 {
		return cmap.segments[0]
	}
	var keyHashHigh int
	if keyHash > math.MaxUint32 {
		keyHashHigh = int(keyHash >> 48)
	} else {
		keyHashHigh = int(keyHash >> 16)
	}
	return cmap.segments[keyHashHigh%cmap.concurrency]
}
