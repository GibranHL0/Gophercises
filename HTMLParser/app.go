package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/GibranHL0/Gophercises/HTMLParser/Link"
)

func main() {
	data, err := os.ReadFile("HTML/ex4.html")

	if err != nil {
		panic(err)
	}

	r := strings.NewReader(string(data))

	links, err := link.Parse(r)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", links)
}
