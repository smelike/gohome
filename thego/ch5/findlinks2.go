package main

import (
	"bytes"
	"fmt"
	"os"

	"net/http"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		/* links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		} */
		fmt.Println(CountWordsAndImages(url))
	}
}

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}

// 黄色波浪线，提示 visit 已经在同一 package main 内重复定义
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	/* 	doc, err := io.ReadAll(resp.Body)
	   	fmt.Println(resp.Body, string(doc), err) */
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("Parsing HTML: %s\n", err)
		return
	}
	words, images, err = countWordsAndImages(doc)
	return
}
func countWordsAndImages(n *html.Node) (words, images int, err error) {
	/*
		words:
		images:
	*/
	// html.Node 转换为 io.Reader
	/*
		scanner := bufio.NewScanner(bufio.NewReader(n))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			w := scanner.Text()
			fmt.Println(w)
		} */
	// html.ElementNode is not html.TextNode
	text := &bytes.Buffer{}
	collectText(n, text)
	fmt.Println(text)
	return
}

func collectText(n *html.Node, buf *bytes.Buffer) {
	if n.Type == html.TextNode {
		buf.WriteString(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		collectText(c, buf)
	}
}
