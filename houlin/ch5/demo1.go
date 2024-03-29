package main

import (
	"errors"
	"io"
	"os"
	"sync"
)

/*
0 案例：文件的读写操作

创建一个文件存放数据，同一时刻可能会由多个 goroutine 进行写操作和读操作。
每次写操作写入若干字节，该若干字节的数据必须作为一独立的数据块存在。
每次读操作从文件读取一个独立完整的数据块，读取的数据块不能重复，且需按顺序读取。
例如，第一个读操作读取了数据块 1，那第二个读操作读取数据块 2，以此类推。
对于读操作是狗可并发进行，并不作要求。即使读是并发进行的，程序应分辨出它们的先后顺序。

0.1 如何考虑【边界情况】？什么是【边界情况】？
*/

/*
1 需求分析：

1)os.File 类型为操作文件系统提供了底层的支持，但并不保证并发操作的安全性。
这里需分别对两类操作(写操作和读操作)做访问控制，所以读写锁比普通的互斥锁更使用。
至于多个读操作要按顺序且不能重复读取，寻求其他辅助手段解决。

2)为实现上述需求，创建了一个接口类型。
*/

// 类型 Data 的声明：被声明为一个 []byte 的别名类型
type Data []byte

/*
WSN - Writing Serial Number
RSN - Reading Serial Number
*/
type DataFile interface {
	// 读取一个数据块
	Read() (rsn int64, d Data, err error)
	// 写入一个数据块
	Write(d Data) (wsn int64, err error)
	// 获取最后读取的数据块的序列号
	RSN() int64
	// 获取最后写入的数据块的序列号
	WSN() int64
	// 获取数据块的长度
	DataLen() uint32
	// 关闭数据文件
	Close()
}

/*
2 接口的实现类型：编写 DataFile 接口的实现类型。
 【接口的实现类型】
 1、多个写操作同时要增加 woffset 字段的值时，会产生竞态条件，需互斥锁 wmutex 以保护。
 2、rmutex 用来消除多个读操作同时增加 roffset 字段值时产生的竞态条件。
 3、数据块的长度在初始化 myDataFile 类型时给定。存储在 dataLen 字段中，与 DataFile 接口中声明的 DataLen 方法是对应的。

*/

type myDataFile struct {
	f       *os.File     // 文件
	fmutex  sync.RWMutex // 文件的读写锁
	woffset int64        // 写操作用到的偏移量
	roffset int64        // 读操作用到的偏移量
	wmutex  sync.Mutex   // 写操作用到的互斥锁
	rmutex  sync.Mutex   // 读操作用到的互斥锁
	dataLen uint32       // 数据块长度
	rcond   *sync.Cond   // 条件变量，为改造边界情况，升级版本 v2 添加的
}

/*
3 创建和初始化 DataFile 类型值的函数: NewDataFile

NewDataFile 函数会返回一个 DataFile 类型的值，但实际会创建并初始化一个 *myDataFile 类型的值，并把它作为其结果值返回。
这样可以通过编译的，因为后者（myDataFile）必须是前者（DataFile）的一个实现类型。

NewDataFile 函数的完整声明如下：
创建 *myDataFile 类型的值时，只需对其中的字段 f 和 dataLen 进行初始化。
因为 woffset 和 roffset 字段的零值都是 0。
至于字段 fmutex/wmutex/rmutex，它们的零值即可用的锁。所以无须对它们进行显式初始化。

把变量 df 的值作为 NewDataFile 函数的第一个结果值，这体现了设计意图。【设计意图】
但要想使 *myDataFile 类型真成为 DataFile 类型的一个实现类型，还需为 *myDataFile 编写 DataFile 接口类型包含的所用方法。

[DataFile] -> [myDataFile]
*/

func NewDataFile(path string, dataLen uint32) (DataFile, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	if dataLen == 0 {
		return nil, errors.New("Invalid data length!")
	}
	df := &myDataFile{f: f, dataLen: dataLen}
	df.rcond = sync.NewCond(df.fmutex.RLocker()) // sync.RLocker() ——> 返回的是读写锁中的读锁
	// 该return df 出会出现 error alert，因为 myDataFile 尚未实现 DataFile 接口
	return df, nil
}

// myDataFile 实现类型，实现 DataFile 接口中的 Read 函数
/*
*myDataFile 类型的 Read 方法，该方法应照如下步骤实现：
0) 版本 0.5
1) 获取并更新 roffset;
(多个读操作不能读取同一个数据块，且应按顺序读取文件中的数据块)
2) 依据 roffset 从文件中读取一块数据;
(读写锁)
3) 把该数据块封装成一个 Data 类型值，并将其作为结果值返回。
4)边界问题：
当有 3 个 goroutine 进行读操作，2 个 goroutine 进行写操作时，
即读操作 goroutine 多于 写操作 goroutine，那么读操作的偏移量最终会赶上写操作。

*/
func (df *myDataFile) Read() (rsn int64, d Data, err error) {
	// 读取并更新读偏移量
	/*

	 */
	var offset int64
	df.rmutex.Lock()
	offset = df.roffset
	df.roffset += int64(df.dataLen)
	df.rmutex.Unlock()

	// 读取一个数据块
	rsn = offset / int64(df.dataLen)
	bytes := make([]byte, df.dataLen)
	/*
		for 语句目的是，在 df.f.ReadAt 方法返回 io.EOF 时，继续尝试获取同一个数据块，直到获取成功为止。
		【注意】：如果在该 for 代码块执行期间一直让读写锁 fmutex 处于读锁定状态，
		那么针对它的写锁定操作将永远不会成功，且相应的 goroutine 也会一直阻塞。
	*/
	df.fmutex.RLock()
	defer df.fmutex.RUnlock()
	for { // version 2: for 死循环监听
		_, err = df.f.ReadAt(bytes, offset)
		if err != nil {
			if err == io.EOF {
				df.rcond.Wait() // 改为使用条件变量，等待通知 <- df.fmutex.RUnlock()
				continue
			}
			// df.fmutex.RUnlock()
			return
		}
		d = bytes
		// df.fmutex.RUnlock()
		return
	}
}

/*
	*myDataFile 类型的 Write 方法，实现步骤：

	1) 获取并更新 woffset
	2) 依据 woffset 写入数据到文件
	3) 返回数据块序列号
*/

func (df *myDataFile) Write(d Data) (wsn int64, err error) {
	var offset int64
	df.wmutex.Lock()
	offset = df.woffset
	df.woffset += int64(df.dataLen)
	df.wmutex.Unlock()

	wsn = offset / int64(df.dataLen) // 获取 write serial number （写入序列号）
	var bytes []byte
	// 写入内容
	if len(d) > int(df.dataLen) { // 写入内容长度大于数据长度时，截取
		bytes = d[0:df.dataLen]
	} else {
		bytes = d
	}
	df.fmutex.Lock()
	defer df.fmutex.Unlock()
	_, err = df.f.Write(bytes)
	df.rcond.Signal() // v2 发送通知，让读操作的 goroutine 进行唤醒等
	return
	// return math.MaxInt64, errors.New("Write data failed!") // just for demo
}

/*
	最后读取的数据块序列号
	步骤：调用互斥锁锁定读偏移量，计算 rsn 返回后，才解锁读偏移量。
*/
func (df *myDataFile) RSN() (rsn int64) {
	// return math.MaxInt64
	df.rmutex.Lock()
	defer df.rmutex.Unlock()
	return df.roffset / int64(df.dataLen) // 读操作的偏移量 【除以】数据块的长度

}

/*
	最后写入的数据块的序列号
	步骤：锁定写偏移量，计算返回 wsn 后，解锁定写偏移量
*/
func (df *myDataFile) WSN() (wsn int64) {
	// return math.MaxInt64
	df.rmutex.Lock()
	defer df.rmutex.Unlock()
	return df.woffset / int64(df.dataLen) // 返回值类型已经被方法签名而声明了
}

/*
	文件的数据块长度
*/
func (df *myDataFile) DataLen() (dl uint32) {
	// dl := df.dataLen // dl := 如此会报错
	dl = df.dataLen
	return
}

/*
	关闭文件
*/
func (df *myDataFile) Close() {
	err := df.f.Close()
	if err != nil {
		return
	}
}
