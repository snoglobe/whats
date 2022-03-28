package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var template = "https://en.wikipedia.org/w/api.php?format=json&action=query&prop=extracts&exintro&explaintext&redirects=1&titles=%s"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: whats <query>")
		os.Exit(1)
	}

	word := strings.Join(os.Args[1:], " ")
	url := fmt.Sprintf(template, url.QueryEscape(word))
	// get the response
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	// get the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	//unmarshal the json
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// get the first page
	pages := data["query"].(map[string]interface{})["pages"]
	page := pages.(map[string]interface{})
	for _, v := range page {
		page := v.(map[string]interface{})
		extract, err := page["extract"].(string)
		if err != true {
			fmt.Printf("%s was not found.\n", word)
			os.Exit(1)
		}
		fmt.Println(extract)
	}
}
