package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	fileBasedPipe()
	inMemorySyncPipe()
}

// 通信：基于文件的通道
func fileBasedPipe() {
	// 创建管道
	reader, writer, err := os.Pipe()
	if err != nil {
		fmt.Printf("Error: Couldn't create the named pipe: %s\n", err)
	}

	// go func() {}() 并发执行?
	go func() {
		output := make([]byte, 100)   // bye 类型
		n, err := reader.Read(output) // 读
		if err != nil {
			fmt.Printf("Error: Couldn't read data from the named pipe: %s\n", err)
		}
		fmt.Printf("Read %d byte(s). [File-based] \n", n)
	}()

	input := make([]byte, 26)
	for i := 65; i <= 90; i++ {
		input[i-65] = byte(i) // 26 个英文小写字母？
	}
	n, err := writer.Write(input) // 写入
	if err != nil {
		fmt.Printf("Error: Couldn't write data to the named pipe: %s\n", err)
	}
	fmt.Printf("Written %d byte(s). [file-based pipe]\n", n)
	time.Sleep(200 * time.Millisecond)
}

// 基于内存的通道
func inMemorySyncPipe() {
	reader, writer := io.Pipe()
	go func() {
		output := make([]byte, 100)
		n, err := reader.Read(output)
		if err != nil {
			fmt.Printf("Error: Couldn't read data from the named pipe: %s\n", err)
		}
		fmt.Printf("Read %d byte(s). [in-memory pipe]\n", n)
	}()
	input := make([]byte, 26)
	for i := 65; i <= 90; i++ {
		input[i-65] = byte(i)
	}
	n, err := writer.Write(input)
	if err != nil {
		fmt.Printf("Error: Couldn't write data to the named pipr: %s\n", err)
	}
	fmt.Printf("Written % byte(s). [in-memory pipe]\n", n)
	time.Sleep(200 * time.Millisecond)
}
