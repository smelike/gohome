package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

func main() {
	var l, d uint64

	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(fmt.Printf("count character by categories: %v\n", err))
		}

		if r == unicode.ReplacementChar && n == 1 {
			continue
		}

		if unicode.IsLetter(r) {
			l++
		} else if unicode.IsDigit(r) {
			d++
		}
	}

	fmt.Printf("type\tcount\n")
	fmt.Println("letters\t", l)
	fmt.Println("digits\t", d)
}

/*
Exercise 4.8: Modify charcount to count letters, digits, and so on in their Unicode categories,
using functions like unicode.IsLetter.

Exercise 4.9: Write a program wordfreq to report the frequency of each word in an input text file.
Call input.Split(bufio.ScanWords) before the first call to Scan to break the input into words instead
of lines.

*/
