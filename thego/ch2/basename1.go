package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	p := "a/b/c.go"
	println(basename(p))
	println(basename2(p))
}

func basename(s string) string {

	var start = time.Now()
	// a/b/c.go

	// len(s) -1 等于是从后往前，反过来寻找最后一个 /
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	fmt.Println("basename1:", time.Since(start).Microseconds())
	return s
}

func basename2(s string) string {
	var since = time.Now()
	slash := strings.LastIndex(s, "/") // -1 if / not found

	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	fmt.Println("basename2:", time.Since(since).Microseconds())
	return s
}
