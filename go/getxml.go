package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Sitemapindex struct {
	Titles []string `xml:"url>loc"`
}

type News struct {
	Keywords  []string `xml:"url>loc"`
	Locations []string `xml:"url>loc"`
	URL       []string `xml:"url"`
}

type NewsMap struct {
	Keyword  string
	Location string
}

func main() {
	// var and map
	var s Sitemapindex
	//var n News
	//	news_map := make(map[string]NewsMap)

	// unmarshal http body using sitemap index
	//resp, _ := http.Get(" ")
	resp, _ := http.Get("https://www.di.se/sitemap/2021-01.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)
	string_body := string(bytes)
	fmt.Println(string_body, "\n")

	// loop to print values
	for idx, Title := range s.Titles {
		fmt.Printf("\n title #%d:\t%s ", idx, Title)
	}
}
