package lib

import "time"

// RawReq\ RawResp defined at loadgen/lib
type Caller interface {
	BuildReq() RawReq
	Call(req []byte, timeoutNS time.Duration) ([]byte, error)
	CheckResp(rawReq RawReq, rawResp RawResp) *CallResult
}
