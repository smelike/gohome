package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var depth int

func main() {

	for _, arg := range os.Args[1:] {

		resp, err := http.Get(arg)

		if err != nil {
			fmt.Println("http get %s: $v", arg, err)
		}

		doc, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Println("parse html failed: %v\n", err)
		}

		forEachNode(doc, startElement, endElement)
	}
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

// The * adverb in %*s prints a string padded with a variable number of spaces.
// The width and the string are provided by the arguments depth*2 and "".
func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
