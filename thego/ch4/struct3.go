package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

// comparable struct types, like other comparable types, used as the key type of a map.

type address struct {
	hostname string
	port     int
}

func main() {
	p := Point{1, 2}
	q := Point{2, 1}

	o := Point{1, 2}

	fmt.Println(p.X == q.X && p.Y == q.Y)
	// == operation compares the corresponding fields of the two structs in order.
	fmt.Println(p == q)

	fmt.Println(p == o)

	hits := make(map[address]int)

	hits[address{"golang.org", 443}]++
	/*
		for _, arg := range os.Args[1:] {
			// hits
		} */
	fmt.Println(hits)
}
