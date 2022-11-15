package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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
	Response bool `json:"exist,omitempty"`
}

type rating struct {
	Source string
	Value  string
}

var n = flag.String("name", "", "a movie name")
var p = flag.String("plot", "short", "movie's plot")

func main() {
	flag.Parse()

	// fmt.Println(flag.Args())
	// fmt.Println(*n)
	if len(strings.TrimSpace(*n)) == 0 {
		log.Fatal(fmt.Println("movie name can not be empty"))
	}
	movie, err := SearchMovie([]string{
		"t=inventing the abbotts",
		"plot=full",
	})
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	/* for _, rate := range movie.Ratings {
		fmt.Printf("%s \t %s\n", rate.Source, rate.Value)
	} */

	p := movie.Poster
	fileName := movie.Title + p[strings.LastIndex(p, "."):]

	err = downloadFile(movie.Poster, fileName)
	if err == nil {
		fmt.Printf("download %s poster successfully\n", movie.Title)
	}
}

func downloadFile(url, fileName string) error {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("image not exist")
		return err
	}
	defer resp.Body.Close()

	f, _ := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
func SearchMovie(terms []string) (*Movie, error) {

	/* query escape:
	http://www.omdbapi.com/?apikey=953924d6&t%3Dinventing+the+abbotts%26plot%3Dfull
	*/
	// q := url.QueryEscape(strings.Join(terms, "&"))
	q := strings.Join(terms, "&")
	fmt.Println(Omdb_URL + q)
	resp, err := http.Get(Omdb_URL + q)
	if err != nil {
		log.Fatal(fmt.Printf("search movies failed: http get %s\n", err))
	}
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
