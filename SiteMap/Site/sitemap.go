package site

import (
	"encoding/xml"
	"strings"

	link "github.com/GibranHL0/Gophercises/HTMLParser/Link"
)

type SiteMap struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	Urls    []Url    `xml:"url"`
}

type Url struct {
	Loc string `xml:"loc"`
}

func ConvertToURLs(links []link.Link, site string) []Url {
	urls := make([]Url, 0)

	for _, link := range links {
		url := link.Href

		if len(url) < 1 || url[0] == '#' || strings.Contains(url, "mailto"){
			continue
		}

		if url[0] == '/'{
			url = site + link.Href
		}

		if contains(urls, url) {
			continue
		}

		tempUrl := Url{
			Loc: url,
		}

		urls = append(urls, tempUrl)
	}

	return urls
}

func CreateSiteMap(site string, urls []Url) SiteMap {
	return SiteMap{
		Xmlns: site,
		Urls:  urls,
	}
}

func CreateXML(sitemap SiteMap) []byte {
	doc, err := xml.MarshalIndent(sitemap, " ", "  ")
	check(err)

	xmldoc := append([]byte(xml.Header), doc...)

	return xmldoc
}

func contains(urls []Url, url string) bool {
	for _, value := range urls {
		if value.Loc == url {
			return true
		}
	}

	return false
}