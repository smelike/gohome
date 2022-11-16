package main

import (
	"fmt"
	"net/http"
)

/*
Exercise 4.14: Create a web server that queries GitHub once
and then allows navigation of the list of bug reports,
milestones, and users.

- a web server that queries Github once
- allow navigation of the list of bug reports, milestones, and users



*/

func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe("localhost:8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path=%q\n", r.URL.Path)
}
