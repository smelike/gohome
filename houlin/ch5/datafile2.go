package v2

import (
	"errors"
	"io"
	"os"
	"sync"
	"sync/atomic"
)

type Data []byte

type DataFile interface {
	Read() (rsn int64, d Data, err error)
	Wirte(d Data) (wsn int64, err error)
	RSN() int64
	WSN() int64
	DataLen() uint32
	Close() error
}

type myDataFile struct {
	f       *os.File
	fmutex  sync.RWMutex
	rcond   *sync.Cond
	woffset int64
	roffset int64
	wmutex  sync.Mutex
	rmutex  sync.Mutex
	dataLen uint32
}

// 实例函数
func NewDataFile(path string, dataLen uint32) (DataFile, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	if dataLen == 0 {
		return nil, errors.New("Invalid data length!")
	}
	df := &myDataFile{f: f, dataLen: dataLen}
	df.rcond = sync.NewCond(df.fmutex.RLocker()) // 条件变量 df.fmutex.RLocker 获取读写锁中的读锁
	// myDataFile 必须实现 DataFile 接口，才能作为 DataFile 类型返回
	return df, nil
}

func (df *myDataFile) Read() (rsn int64, d Data, err error) {
	var offset int64
	/*
		// 使用互斥锁的案例代码
		df.rmutex.Lock()
		offset = df.roffset
		df.roffset += int64(df.dataLen)
		df.rmutex.Unlock()
	*/
	// 使用原子函数
	// 字段 roffset 和变量 offset 都是 int64 类型的，后者代表了前者的旧值。
	// 而字段 roffset 的新值即其旧值与 dataLen 字段值的和。
	for {
		// 在 32 位计算架构的计算机上写入一个 64 位的整数，会存在并发安全方面的隐患
		// offset = df.roffset
		offset = atomic.LoadInt64(&df.roffset)
		if atomic.CompareAndSwapInt64(&df.roffset, offset, (offset + int64(df.dataLen))) {
			break
		}
	}
	rsn = offset / int64(df.dataLen)
	bytes := make([]byte, df.dataLen)
	df.fmutex.RLock()
	defer df.fmutex.RUnlock()
	for {
		_, err = df.f.ReadAt(bytes, offset) // use := will have error alert why??
		if err != nil {
			if err == io.EOF {
				df.rcond.Wait()
				continue
			}
			return
		}
		d = bytes
		return
	}
}

func (df *myDataFile) Write(d Data) (wsn int64, err error) {

	// offset 是一个函数内的局部变量
	var offset int64
	/*
		df.wmutex.Lock()
		offset = df.woffset
		df.woffset += int64(df.dataLen)
		df.wmutex.Unlock()
	*/
	for {
		offset = atomic.LoadInt64(&df.woffset)
		if atomic.CompareAndSwapInt64(&df.woffset, offset, (offset + int64(df.dataLen))) {
			break // 更新更成功则跳出循环监听
		}
	}
	// 写入
	wsn = offset / int64(df.dataLen)
	var bytes []byte
	if len(d) > int(df.dataLen) {
		bytes = d[0:df.dataLen]
	} else {
		bytes = d
	}
	df.fmutex.Lock()
	defer df.fmutex.Unlock()
	_, err = df.f.Write(bytes)
	df.rcond.Signal()
	return
}

func (df *myDataFile) RSN() (rsn int64) {
	// df.rmutex.Lock()
	// defer df.rmutex.Unlock()
	offset := atomic.LoadInt64(&df.roffset)
	return offset / int64(df.dataLen)
}

func (df *myDataFile) WSN() (wsn int64) {
	// df.wmutex.Lock()
	// defer df.wmutex.Unlock()
	// return df.woffset / int64(df.dataLen)
	offset := atomic.LoadInt64(&df.woffset)
	return offset / int64(df.dataLen)
}

func (df *myDataFile) DataLen() uint32 {
	return df.dataLen
}

func (df *myDataFile) Close() error {
	if df.f == nil {
		return nil
	}
	return df.f.Close()
}
