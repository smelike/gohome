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

func getType3(content string) (ty string) {
	switch lang := strings.TrimSpace(content); lang {
	case "Ruby":
		fallthrough
	case "Python":
		ty = "An interpreted Language"
	case "C", "Java", "Go":
		ty = "A compiled language"
	default:
		ty = "Unknow language"
	}
	return
}

func getType4(v interface{}) (ty string) {
	// var v interface{}
	switch v.(type) {
	case string:
		ty = fmt.Sprintf("The string is '%s.'\n", v.(string))
	case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
		ty = fmt.Sprintf("The integer is %d. \n", v)
	default:
		ty = fmt.Sprintf("Unsupported value. (type=%T)\n", v)
	}
	return
}

func getType5(v interface{}) (ty string) {
	switch i := v.(type) {
	case string:
		ty = fmt.Sprintf("The string is '%s'. \n", i)
	case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
		ty = fmt.Sprintf("The integer is %d.\n", i)
	default:
		ty = fmt.Sprintf("Unsupported value. (type=%T)\n", i)
	}
	return
}

var myInt int = 1200
var myName string = "Who is the path?"

func main() {
	ty := getType("Java")
	fmt.Println(ty)

	fmt.Println("-----getType2-------------")
	ty = getType2("  Go    ")
	fmt.Println(ty)

	fmt.Println("-----getType3------------")
	ty = getType3("Python")
	fmt.Println(ty)
	fmt.Println("----------getType4-----------")
	fmt.Println(myInt)
	ty = getType4(myInt)
	fmt.Println(ty)

	fmt.Println("----------getType5----------")
	ty = getType5(myName)
	fmt.Println(ty)
}
