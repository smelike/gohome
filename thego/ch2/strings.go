package main

import "fmt"

func main() {
	s := "hello, world"

	fmt.Println(len(s))
	// 32 代表空格？
	fmt.Println(s[6])

	const GoUsage = `Go is a tool for managing Go source code.
		Usage:
			go command [arguments] \v \\\
		
	....`

	fmt.Println(GoUsage)
}
