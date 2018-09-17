package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// SitemapIndex - This is for all sitemap XML tags
type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`
}

// Location - This is for all loc XML tags
type Location struct {
	Loc string `xml:"loc"`
}

func (location Location) String() string {
	return fmt.Sprintf(location.Loc)
}

func main() {
	response, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	responseBody, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()

	var sitemap SitemapIndex
	xml.Unmarshal(responseBody, &sitemap)

	// fmt.Println(sitemap.Locations)
	for index, siteLocation := range sitemap.Locations {
		fmt.Println(index, ":", siteLocation)
	}
}
