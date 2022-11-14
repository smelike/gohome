package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	Exercise 4.12: The popular web comic xkcd has a JSON interface. For example, a request
	to https://xkcd.com/571/info.0.json produces a detailed description of comic 571, one of
	many favorites. Download each URL (once!) and build an offline index.

	Write a tool xkcd that, using this index, prints the URL and transcript of each that matches
	a search term provide on the command line.

	获取请求状态码为 200 index list，写入到文件中

	https://xkcd.com/571/info.0.json

	获取 transcript of each comic

	获取命令行中的搜索词 os.Args[1:]
	matches a search term provide on the command line


	打印匹配搜索的 each comic 的 URL 和 transcript
*/

type MatchResult []Comic

type Comic struct {
	Number     int `json:"num"`
	Title      string
	Transcript string
	Img        string
	Alt        string
}

func main() {
	result, err := GenerateIndex()
	// fmt.Println(len(*result), err)
	if err != nil {
		fmt.Printf("Generate comic index failed: %s", err)
	}
	for _, item := range *result {
		// fmt.Println(json.MarshalIndent(item, "", "~"))
		fmt.Printf("%d - %s\n", item.Number, item.Title)
		// break
	}
}

func GenerateIndex() (*MatchResult, error) {

	var result MatchResult
	var comic Comic
	for i := 1; i < 10; i++ {
		url := "https://xkcd.com/" + fmt.Sprint(i) + "/info.0.json"
		fmt.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			continue
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			return nil, fmt.Errorf("request failed: %s ", url)
		}
		if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
			resp.Body.Close()
			continue
		}
		// 如果使用无限循环 for {}, 何时代表结束循环呢？
		result = append(result, comic)
		resp.Body.Close()
	}
	return &result, nil
}
