package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
		/*
			A goroutine is [a concurrent function execution].
			A channel is a communication mechanism that allows one goroutine
			to pass [values of a specified type] to another goroutine.
			The function main runs in a goroutine and the go statement creates additional goroutines.

		*/
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed \n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	// 如何判断是否有前缀 https 和 http
	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Error: %v", err)
		return
	}
	// the number of bytes
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	// time.Since(start) == time.Now().sub(t).
	secs := time.Since(start).Seconds()
	// 往通道发送摘要的字符串
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

/*
When one goroutine attempts a send or receive on a channel, it blocks until
another goroutine attempts the corresponding receive or send operation, at
which point the value is transferred and both goroutine proceed.

In this example, each fetch sends a value (ch <- expression) on the channel ch,
and main receives all of them (<-ch).

Having main do all the printing ensures that output from each goroutine is processed as
a unit, with no danger of interleaving if two goroutines finish at the same time.


*/

/*
Exercise 1.10: Find a web site that produces a large amount of data. Investigate caching by
running fetchall twice in succession to see whether the reported time changes much.
Do you get the same content each time? Modify fetchall to print its output to a file so it can
be examined.

[large amount of data - the reported time changes]

Exercise 1.11: Try fetchall with longer argument lists, such as samples from the top million
web sites available at alexa.com. How does the program behave if a web site just
doesn't respond?

[longer argument lists - the top million web site]

*/
