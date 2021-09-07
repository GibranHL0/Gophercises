package main

import (
	"fmt"

	site "github.com/GibranHL0/Gophercises/SiteMap/Site"
)

func MySolution() {
	url := "https://www.calhoun.io"

	html := site.GetHTML(url)

	links := site.GetLinks(html)

	urls := site.ConvertToURLs(links, url)

	sitemap := site.CreateSiteMap(url, urls)

	xmldoc := site.CreateXML(sitemap)

	fmt.Println(string(xmldoc))

	// fmt.Printf("%+v\n", links)
}

func GopherciseSolution() {
	domain := "https://www.calhoun.io"
	maxDepth := 10

	fmt.Println("Domain to be checked: ", domain)

	urls := site.ConvertToUrls(site.Bfs(domain, maxDepth))

	sitemap := site.CreateSiteMap(domain, urls)

	xmldoc := site.CreateXML(sitemap)

	fmt.Println(string(xmldoc))
}

func main() {
	GopherciseSolution()
}
