package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
report the frequency of each word in an input text file.


 step1: 打开文件
 step2: 读取文件内容，注意处理大文件所需采取的缓存
 step3: 内容的字符拆分处理，并计数

*/

func main() {
	var counts = make(map[string]int)
	// fmt.Println(counts == nil)
	for _, filename := range os.Args[1:] {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		scan := bufio.NewScanner(f)
		scan.Split(bufio.ScanWords)
		for scan.Scan() {
			// 消除 , . ' "" 等等
			w := strings.Trim(scan.Text(), ",|.")
			// store to a nil map will panic
			counts[strings.ToLower(w)]++
		}
	}

	fmt.Printf("Key\t\t\tFreq\n")
	for k, c := range counts {
		fmt.Printf("%s\t\t\t%d\n", k, c)
	}
}
