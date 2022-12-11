package main

import (
	"fmt"
	"links"
	"log"
	"os"
)

func main() {
	/* for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("http get failed: %s", err)
		}
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("getting %s: %s\n", url, err)
		}
	} */
	breadthFirst(crawl, os.Args[1:])
}

/*
Append returns the updated slice. It is therefore necessary to store the
result of append, often in the variable holding the slice itself:
	slice = append(slice, elem1, elem2)
	slice = append(slice, anotherSlice...)
As a special case, it is legal to append a string to a byte slice, like this:
	slice = append([]byte("hello "), "world"...)
*/
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
