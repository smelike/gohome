package main

import (
	"fmt"
	"time"
)

// A struct is a collection of fields.
type Vertext struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertext{int(time.Second), 2})
}
