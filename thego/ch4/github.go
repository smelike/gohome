package github

import (
	"encoding/json"
	"fmt"
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
	TotalCount int `json:"total_count"`
	Items      []*Issue
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
/* func main() {
	var terms []string
	terms = append(terms, "repo:MixinNetwork/mixin", "type:issue")
	// fmt.Printf("%v\t%[1]T", terms)
	result, err := SearchIssues(terms)
	// fmt.Printf("%v\t%s", result, err)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues: \n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d [%9.9s] %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
} */

func SearchIssues(terms []string) (*IssueSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	// q := strings.Join(terms, "&")
	fmt.Println(IssuesURL + "?q=" + q)
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
