package main

import "fmt"

func main() {
	var chars []string = []string{"abc", "nm", "", "world"}
	data := []string{"one", "", "three", "four"}
	fmt.Println(nonempty(chars))
	fmt.Println(chars)
	fmt.Println(nonempty(data))
	fmt.Println(data)
}

func nonempty(strings []string) []string {
	i := 0 // 保存使用的索引
	fmt.Printf("start=%v\n", strings)
	for _, s := range strings {
		if s != "" {
			fmt.Printf("if==%v\n", strings)
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s) // assign and overite slice out
		}
	}
	return out
}
