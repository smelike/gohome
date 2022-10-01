package loadgen

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math"
	"sync/atomic"
	"time"

	lib "./lib"
)

// var logger = log.DLogger()

// 结构体 myGenerator
type myGenerator struct {
	caller      lib.Caller           // 调用器
	timeoutNS   time.Duration        // 处理超时时间，单位：纳秒
	lps         uint32               // 每秒载荷量
	durationNS  time.Duration        // 负载持续时间，单位：纳秒
	concurrency uint32               // 载荷并发数
	tickets     lib.GoTickets        // Goroutine票池
	ctx         context.Context      // 上下文
	cancelFunc  context.CancelFunc   // 取消函数
	callCount   int64                // 调用技术
	status      uint32               // 状态
	resultCh    chan *lib.CallResult // 调用结果通道
}

// 新建载荷发生器
func NewGenerator(pset ParamSet) (lib.Generator, error) {
	if err := pset.Check(); err != nil {
		return nil, err
	}
	gen := &myGenerator{
		caller:     pset.Caller,
		timeoutNS:  pset.TimeoutNS,
		lps:        pset.LPS,
		durationNS: pset.DurationNS,
		status:     lib.STATUS_ORIGINAL,
		result:     pset.ResultCh,
	}
	if err := gen.init(); err != nil {
		return nil, err
	}
	return gen, nil
}

// 初始化载荷发生器
func (gen *myGenerator) init() error {
	var buf bytes.Buffer
	buf.WriteString("Initializing the load generator...")
	// 载荷并发量 约等于 载荷的响应超时时间 / 载荷的发送间隔时间
	var total64 = int64(gen.timeoutNS)/int64(1e9/gen.lps) + 1
	if total64 > math.MaxInt32 {
		total64 = math.MaxInt32
	}
	gen.concurrency = uint32(total64)
	tickets, err := lib.NewGoTickets(gen.concurrency)
	if err != nil {
		return err
	}
	gen.tickets = tickets
	buf.WriteString(fmt.Sprintf("Done. (concurrency = %d)", gen.concurrency))
	// 日志记录
	return nil
}

// 向载荷承受方发起一次调用
func (gen *myGenerator) callone(rawReq *lib.RawReq) *lib.RawResp {
	atomic.AddInt64(&gen.callCount, 1)
	if rawReq == nil {
		return &lib.RawResp{ID: -1, Err: errors.New("Invalid raw request.")}
	}
	start := time.Now().UnixNano()
	resp, err := gen.caller.Call(rawReq.Req, gen.timeoutNS)
	end := time.Now().UnixNano()
	elapsedTime := time.Duration(end - start)
	var rawResp lib.RawResp
	if err != nil {
		errMsg := fmt.Sprintf("Sync Call Error: %s.", err)
		rawResp = lib.RawResp{
			ID:     rawReq.ID,
			Err:    errors.New(errMsg),
			Elapse: elapsedTime}
	} else {
		rawResp = lib.RawResp{
			ID:     rawReq.ID,
			Resp:   resp,
			Elapse: elapsedTime}
	}
	return &rawResp
}
