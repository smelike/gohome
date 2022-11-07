package main

import (
	"fmt"
	"sort"
)

func main() {

	// var names []string // a slice of string
	ages := map[string]int{
		"alice":   33,
		"charlie": 34,
		"intel":   16,
		"dell":    21,
		"apple":   20,
	}
	// allocate an array of the required size up front
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}

	sort.Strings(names) // 字符串排序

	// 以排序好的名字，循环获取年龄
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	var books map[string]int
	fmt.Println(books == nil)
	fmt.Println(len(books) == 0)
	books = map[string]int{}
	books["amazon"] = 12
}
