package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

/*
first define the necessary types and constants

*/

const IssuesURL = "https://api.github.com/search/issues"

// 搜索结果
type IssueSearchResult struct {
	TotalCount int      `json:"total_count"`
	Items      []*Issue //  A slice: the value type is *Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

/*
 repo:mixin -> search query failed: 422 Unprocessable Entity
*/
func main() {
	var terms []string = []string{
		"repo:MixinNetwork/mixin",
		"type:issue",
	}
	// terms = append(terms, "repo:MixinNetwork/mixin", "type:issue")
	// fmt.Printf("%v\t%[1]T", terms)
	result, err := SearchIssues(terms)
	// fmt.Printf("%v\t%s", result, err)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues: \n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d [%9.9s] %.55s %v\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
}

func SearchIssues(terms []string) (*IssueSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	// q := strings.Join(terms, "&")
	// fmt.Println(IssuesURL + "?q=" + q)
	resp, err := http.Get(IssuesURL + "?q=" + q)
	// network request
	if err != nil {
		return nil, err
	}
	// request response
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssueSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

/*
	Exercise 4.10: Modify issues to report the results in age categories, say
	less than a month old, less than a year old, and more than a year old.

	Exercise 4.11: Build a tool that lets users create, read, update, and delete
	GitHub issues from the command line, invoking their preferred text editor when
	substantial text input is required.

	Exercise 4.12: The popular web comic xkcd has a JSON interface. For example, a request
	to https://xkcd.com/571/info.0.json produces a detailed description of comic 571, one of
	many favorites. Download each URL (once!) and build an offline index.
	Write a tool xkcd that, using this index, prints the URL and transcript of each that matches
	a search term provide on the command line.


	Exercise 4.13: The JSON-based web service of the Open Movie Database lets you search
	https://omdbapi.com/ for a movie by name and download its poster image. Write a tool poster
	that downloads the poster image for the movie named on the command line.

*/