package main

import "sync"

var (
	mu      sync.Mutex // guards mapping
	mapping = make(map[string]string)
)

func Loopup(key string) string {
	mu.Lock()
	v := mapping[key]
	mu.Unlock()
	return v
}

/*
version 2
*/

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string), // initialize the map
}

func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}
