package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// SitemapIndex - All sitemap XML tags
type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

// News - All news items
type News struct {
	Titles    []string `xml:"url>news>title"`
	KeyWords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

func main() {
	var sitemap SitemapIndex
	var news News

	response, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	responseBody, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()

	xml.Unmarshal(responseBody, &sitemap)

	for index, link := range sitemap.Locations {
		fmt.Println(index, ":", link)

		response, _ := http.Get(link)
		responseBody, _ := ioutil.ReadAll(response.Body)
		xml.Unmarshal(responseBody, &news)

		fmt.Println(news.Titles)
	}
}
