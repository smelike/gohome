package main

import (
	"errors"
	"math"
	"os"
	"sync"
)

/*
0 案例：
创建一个文件存放数据，同一时刻可能会由多个 goroutine 进行写操作和读操作。
每次写操作写入若干字节，该若干字节的数据必须作为一独立的数据块存在。
每次读操作从文件读取一个独立完整的数据块，读取的数据块不能重复，且需按顺序读取。
例如，第一个读操作读取了数据块 1，那第二个读操作读取数据块 2，以此类推。
对于读操作是狗可并发进行，并不作要求。即使读是并发进行的，程序应分辨出它们的先后顺序。

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
	// 该return df 出会出现 error alert，因为 myDataFile 尚未实现 DataFile 接口
	return df, nil
}

// myDataFile 实现类型，实现 DataFile 接口中的 Read 函数
/*
*myDataFile 类型的 Read 方法，该方法应照如下步骤实现：

1) 获取并更新 roffset;
2) 依据 roffset 从文件中读取一块数据;
3) 把该数据块封装成一个 Data 类型值，并将其作为结果值返回。
*/
func (df *myDataFile) Read() (rsn int64, d Data, err error) {
	d := &Data{}
	return math.MaxInt64, d, errors.New("Read data failed!")
}

/*
	*myDataFile 类型的 Write 方法，实现步骤：

	1) 获取并更行 woffset
	2) 依据 woffset 写入数据到文件
	3) 返回数据块序列号
*/

func (df *myDataFile) Write(d Data) (wsn int64, err error) {
	return math.MaxInt64, errors.New("Write data failed!")
}

/*
	最后读取的数据块序列号
*/
func (df *myDataFile) RSN() (rsn int64) {
	return math.MaxInt64
}

/*
	最后写入的数据块的序列号
*/
func (df *myDataFile) WSN() (wsn int64) {
	return math.MaxInt64
}

/*
	文件的数据块长度
*/
func (df *myDataFile) DataLen() (dl uint32) {
	return math.MaxUint32
}

/*
	关闭文件
*/
func (df *myDataFile) Close() {

}
