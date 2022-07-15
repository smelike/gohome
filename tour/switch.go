package main

import (
	"fmt"
	"runtime"
)

//  switch cases need not be constants [cases], [values]and the values invloved need not be integers
func switch_main() {
	fmt.Print("Go runs on ")
	// fmt.Println() -> mean print line

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}
