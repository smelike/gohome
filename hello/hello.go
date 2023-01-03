package main

import (
	"fmt"
	"time"

	"example.com/hello/morestrings"
	"github.com/google/go-cmp/cmp"
)

/* import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hello("Gladys")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
*/

func main() {
	fmt.Println("Hello, world.", time.Now())

	fmt.Println(morestrings.ReverseRunes("!oG, olleH"))
	fmt.Println(cmp.Diff("hello world", "hello go"))
}
