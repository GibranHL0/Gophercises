package main

import (
	"fmt"

	site "github.com/GibranHL0/Gophercises/SiteMap/Site"
)

func main() {
	url := "https://www.calhoun.io"

	html := site.GetHTML(url)

	links := site.GetLinks(html)

	urls := site.ConvertToURLs(links, url)

	sitemap := site.CreateSiteMap(url, urls)

	xmldoc := site.CreateXML(sitemap)

	fmt.Println(string(xmldoc))

	// fmt.Printf("%+v\n", links)
}
