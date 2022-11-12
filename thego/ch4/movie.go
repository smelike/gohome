package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// struct Movie
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
	rates  []float32
}

// slices []Movie
var movies = []Movie{
	{Title: "Casablance", Year: 19242, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func main() {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	// json.MarshalIndent(movies, prefix, indentation)

	/* data, err = json.MarshalIndent(movies, "", "~")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data) */

	// the inverse operation to marshaling, decoding JSON and populating a Go data structure.

	var titles []struct {
		Title string
		Color bool
	}
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles)
}
