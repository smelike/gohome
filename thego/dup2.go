package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		处理标准输入 os.Stdin 或输入的是文件路径
	*/
	counts := make(map[string]int)
	files := os.Args[1:]

	// 参数为零时，则转为 os 标准输入
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		// 考虑产生竞态条件的情况？
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// 统计
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
