package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	// 获取文件名
	for _, filename := range os.Args[1:] {
		// 打开文件
		data, err := ioutil.ReadFile(filename)
		// 处理异常
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		// 统计行数
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	// 打印统计结果
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
