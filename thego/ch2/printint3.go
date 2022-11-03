package main

import (
	"bytes"
	"fmt"
)

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

/*
	When appending the UTF-8 encoding of an arbitrary run to a bytes.Buffer, it's
	best to use bytes.Buffer's WriteRune method, but WriteByte is fine for ASCII characters
	such as '[' and ']'
*/

func main() {
	fmt.Println(intsToString([]int{12, 34, 90}))
}
