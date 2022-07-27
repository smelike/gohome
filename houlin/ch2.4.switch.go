package main

import (
	"fmt"
	"strings"
)

var content string

func getType(content string) (ty string) {
	switch content {
	default:
		ty = "Unknown language"
	case "Python":
		ty = "An interpreted Language"
	case "Go":
		ty = "A compiled language"
	}
	ty = fmt.Sprintf("%v is %v", content, ty)
	return
}

func getType2(content string) (ty string) {
	switch lang := strings.TrimSpace(content); lang {
	default:
		ty = "Unknown language"
	case "Python":
		ty = "An interpreted language"
	case "Go":
		ty = "A compiled language"
	}
	ty = fmt.Sprintf("%v is %v", strings.TrimSpace(content), ty)
	return
}
func main() {
	ty := getType("Java")
	fmt.Println(ty)

	ty = getType2("  Go    ")
	fmt.Println(ty)
}
