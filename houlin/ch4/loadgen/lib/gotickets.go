package lib

import (
	"errors"
	"fmt"
)

// 接口
type GoTickets interface {
	Take()
	Return()
	Active() bool
	Total() uint32
	Remainder() uint32
}

// 实现
type myGoTickets struct {
	total    uint32        // 票的总数
	ticketCh chan struct{} // 票的容器
	active   bool          // 票池是否已被激活
}

// 新建一个 Goroutine 票池
func NewGoTickets(total uint32) (GoTickets, error) {
	gt := myGoTickets{}
	if !gt.init(total) {
		errMsg := fmt.Sprintf("The goroutine ticket pool can NOT be initialized! (total = %d)\n", total)
		return nil, errors.New(errMsg)
	}
	return &gt, nil
}

// 继承结构体 myGoTickets 拓展方法 init
func (gt *myGoTickets) init(total uint32) bool {
	if gt.active {
		return false
	}
	// 票数的大小
	if total == 0 {
		return false
	}
	// 创建以结构体为值的通道
	ch := make(chan struct{}, total)
	// uint32 转换为 int
	n := int(total)
	for i := 0; i < n; i++ {
		ch <- struct{}{} // 初始化 ch
	}
	gt.ticketCh = ch
	gt.total = total
	gt.active = true
	return true
}

func (gt *myGoTickets) Take() {
	<-gt.ticketCh // 从 myGoTickets 的 ticketCh 取得一个
}

func (gt *myGoTickets) Return() {
	gt.ticketCh <- struct{}{}
}

func (gt *myGoTickets) Active() bool {
	return gt.active
}

func (gt *myGoTickets) Total() uint32 {
	return gt.total
}

func (gt *myGoTickets) Remainder() uint32 {
	return uint32(len(gt.ticketCh))
}
