package lib

import "time"

type RawReq struct {
	ID  int64
	Req []byte
}

type RawResp struct {
	ID     int64
	Resp   []byte
	Err    error
	Elapse time.Duration
}

type RetCode int

// 保留 1 ~ 1000 给载荷承受方使用

const (
	RET_CODE_SUCCESS              = 0
	RET_CODE_WARNING_CALL_TIMEOUT = 1001
	RET_CODE_ERROR_CALL           = 2001
	RET_CODE_ERROR_RESPONSE       = 2002
	RET_CODE_ERROR_CALEE          = 2003
	RET_CODE_FATAL_CALL           = 3001
)

// 依据结果代码返回响应的文字解释
func GetCodePlain(code RetCode) string {
	var codePlain string
	switch code {
	case RET_CODE_SUCCESS:
		codePlain = "Success"
	case RET_CODE_WARNING_CALL_TIMEOUT:
		codePlain = "Call Timeout Warning"
	case RET_CODE_ERROR_CALL:
		codePlain = "Call Error"
	case RET_CODE_ERROR_RESPONSE:
		codePlain = "Response Error"
	case RET_CODE_ERROR_CALEE:
		codePlain = "Callee Error"
	case RET_CODE_FATAL_CALL:
		codePlain = "Call Fatal Error"
	default:
		codePlain = "Unknown result code"
	}
	return codePlain
}

type CallResult struct {
	ID     int64         // ID
	Req    RawReq        // 原生请求
	Resp   RawResp       // 原生响应
	Code   RetCode       // 响应代码
	Msg    string        // 结果成因的简述
	Elapse time.Duration // 耗时
}

// 载荷发生器状态的常量
const (
	STATUS_ORIGINAL uint32 = 0 // 原始
	STATUS_STARTING uint32 = 1 // 启动
	STATUS_STARTED  uint32 = 2 // 已启动
	STATUS_STOPPING uint32 = 3 // 正在停止
	STATUS_STOPPED  uint32 = 4 // 已停止
)

// 载荷发生器的接口：行为定义
type Generator interface {
	Start() bool
	Stop() bool
	Status() uint32
	CallCount() int64
}
