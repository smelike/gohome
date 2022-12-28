package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (t rot13Reader) Read(b []byte) (int, error) {
	n, err := t.r.Read(b)
	for x := range b {
		c := b[x]
		if (c >= 'a' && c <= 'm') || (c >= 'A' && c <= 'M') {
			b[x] += 13
		} else if (c >= 'n' && c <= 'z') || (c >= 'N' && c <= 'Z') {
			b[x] -= 13
		}
	}
	return n, err
	/*
		A B C D E F G H I J K L M

		N O P Q R S T U V W X Y Z
	*/
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
