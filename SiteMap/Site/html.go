package site

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	link "github.com/GibranHL0/Gophercises/HTMLParser/Link"
)

func GetHTML(url string) []byte {
	resp, err := http.Get(url)
	check(err)

	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	check(err)

	return html
}

func GetLinks(html []byte) []link.Link {
	reader := strings.NewReader(string(html))

	links, err := link.Parse(reader)
	check(err)

	return links
}

func hrefs(r io.Reader, base string) []string {
	links, err := link.Parse(r)
	check(err)

	var urls []string

	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			urls = append(urls, base+l.Href)

		case strings.HasPrefix(l.Href, "http"):
			urls = append(urls, l.Href)
		}
	}

	return urls
}

func getPages(urlStr string) []string {
	// Get the HTML from the domain
	resp, err := http.Get(urlStr)
	check(err)

	defer resp.Body.Close()

	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()

	return filter(hrefs(resp.Body, base), withPrefix(base))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
