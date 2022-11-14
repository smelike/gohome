package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

/*
Exercise 4.13:
The JSON-based web service of the Open Movie Database lets you search
https://omdbapi.com/ for a movie by name and download its poster image.

Write a tool poster that downloads the poster image for the movie named on the command line.

http://www.omdbapi.com/?t=inventing+the+abbotts
http://www.omdbapi.com/?t=inventing+the+abbotts&plot=full&y=

{
	"Title":"Inventing the Abbotts","Year":"1997","Rated":"R",
	"Released":"04 Apr 1997","Runtime":"110 min","Genre":"Drama, Romance",
	"Director":"Pat O'Connor","Writer":"Sue Miller, Ken Hixon",
	"Actors":"Liv Tyler, Jennifer Connelly, Joaquin Phoenix",
	"Plot":"Two working class brothers court three wealthy and beautiful sisters in a small Illinois town.",
	"Language":"English","Country":"United States","Awards":"N/A",
	"Poster":"https://m.media-amazon.com/images/M/MV5BZTY0YWM0MDgtYjE2Yi00ZWExLWE0M2MtN2U3YzRjMWNjMjExXkEyXkFqcGdeQXVyNTM0NTU5Mg@@._V1_SX300.jpg",
	"Ratings":[
		{"Source":"Internet Movie Database","Value":"6.4/10"},
		{"Source":"Rotten Tomatoes","Value":"33%"},
		{"Source":"Metacritic","Value":"49/100"}
	],
	"Metascore":"49","imdbRating":"6.4","imdbVotes":"13,108",
	"imdbID":"tt0119381","Type":"movie","DVD":"13 Mar 2001",
	"BoxOffice":"$5,936,344","Production":"N/A",
	"Website":"N/A","Response":"True"
}
*/

/*
	取得电影名
	os.Args[1:] -name -y -plot(default: short, full)

 	构建搜索请求链接，发出请求，获取响应，从响应中取得海报的图片链接
	string.Join(term, "&") url.QueryEscape(), http.Get(), json.NewDecoder().Decode()
 	下载图片

	http://www.omdbapi.com/?i=tt3896198&apikey=953924d6

*/

const Omdb_URL = "http://www.omdbapi.com/?apikey=953924d6&"

type Movie struct {
	Title    string
	Poster   string
	Genre    string
	Director string
	Language string
	Country  string
	Actors   string
	Plot     string
	Ratings  []rating
}

type rating struct {
	Source string
	Value  string
}

func main() {

	movie, err := SearchMovie([]string{
		"t=inventing the abbotts",
		"plot=full",
	})
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	fmt.Println(movie.Title, movie.Poster)
}

func SearchMovie(terms []string) (*Movie, error) {

	q := url.QueryEscape(strings.Join(terms, "&"))

	resp, err := http.Get(Omdb_URL + q)
	if err != nil {
		log.Fatal(fmt.Printf("search movies failed: http get %s\n", err))
	}
	// fmt.Println(resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Search movies response error")
	}

	var movie Movie
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &movie, nil
}
