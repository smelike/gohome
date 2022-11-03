package main

import (
	"bytes"
	"fmt"
	"strings"
)

/*
	每 3 位数字前加逗号
*/
func main() {
	var s string = "12345678"

	fmt.Println(comma(s))

	var f string = "555999.66690"
	fmt.Println(comma2(f))

	fmt.Println(comma3(s))
}

/*
The argument to comma is a string. If its length is less than or equal to 3,
no comma is necessary. Otherwise, comma calls itself recursively with a substring
consisting of all but the last three characters, and appends a comma and the last
three characters to the result of the recursive call.
*/
// 递归调用函数
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	// return 需要不断处理的字符串
	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma2(s string) string {
	i := strings.Index(s, ".")
	// 如果是浮点数
	if i != -1 {
		// s := s[:i]
		return comma(s[:i]) + "." + comma(s[i+1:])
	} else {
		return comma(s)
	}
}

// use bytes.Buffer
func comma3(s string) string {
	var buf bytes.Buffer

	// buf.WriteString(s[])
	for i := len(s) - 1; i > 0; i-- {
		if buf.Len()%3 == 0 {
			buf.WriteString(", ")
		}
		buf.WriteByte(s[i])
		// fmt.Fprintf(&buf, "%d", s[i])
	}
	return buf.String()
}

/*
Exercise 3.10： Write a non-recursive version of comma, using bytes.Buffer instead of
string concatenation.

Exercise 3.11: Enhance comma so that it deals correctly with floating-point numbers and
an optional.

Exercise 3.12: Write a function that reports whether two strings are anagrams of each other,
that is, they contain the same letters in a different order.

[anagram: a word, phrase, or name formed by rearranging the letters of another, such as spar, formed from rasp.]
*/
