package site

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/GibranHL0/Gophercises/HTMLParser/Link"
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}
