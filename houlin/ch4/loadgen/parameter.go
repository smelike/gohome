package loadgen

import (
	"context"
	"time"
)

// 去掉 logger

type myGenerator struct {
	caller      lib.caller    // 调用器
	timeoutNS   time.Duration // 处理超时时间，单位：纳秒
	lps         uint32
	durationNS  time.Duration
	concurrency uint32
	tickets     lib.GoTickets
	ctx         context.Context
	cancelFunc  context.CancelFunc
	callCount   int64
	status      uint32
	resultCh    chan *lib.CallResult
}
