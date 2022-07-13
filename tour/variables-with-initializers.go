package main

import (
	"fmt"
	"math/rand"
)

var i, j int = rand.Intn(100), rand.Intn(366)

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
