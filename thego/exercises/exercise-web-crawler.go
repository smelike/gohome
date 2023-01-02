package main

import (
	"fmt"
	"sync"
)

/*
	1/ map 不是并发安全的
	2/ func main() 中的 goroutine 执行结束了，则其他所有的 goroutine 都将会结束，即因此而无法运行。
	// var fetchedUrl map[string]bool
*/

type Cache struct {
	url map[string]bool
	mu  sync.Mutex
}

func (c *Cache) get(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.url[key]
}

func (c *Cache) set(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.url[key] = true
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

var cache = &Cache{url: make(map[string]bool)}
var wg sync.WaitGroup

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	defer wg.Done()
	if depth <= 0 {
		return
	}
	if cache.get(url) {
		return
	}
	cache.set(url)
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth-1, fetcher)
	}
	return
}

/*
Add adds delta, which may be negative, to the WaitGroup counter. If the counter becomes zero, all goroutines blocked on Wait are released. If the counter goes negative, Add panics.

Note that calls with a positive delta that occur when the counter is zero must happen before a Wait.
Calls with a negative delta, or calls with a positive delta that start when the counter is greater than zero, may happen at any time.
Typically this means the calls to Add should execute before the statement creating the goroutine or other event to be waited for.
If a WaitGroup is reused to wait for several independent sets of events, new Add calls must happen after all previous Wait calls have returned. See the WaitGroup example.
*/
func main() {
	wg.Add(1)
	Crawl("https://golang.org/", 4, fetcher)
	wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

// 实现了 type Fetch interface 的接口
func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
