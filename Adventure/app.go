package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Arc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func storyHandler(intro Arc, templ *template.Template) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		templ.Execute(rw, intro)
	}
}

func defaultMux(stories map[string]Arc, templ *template.Template) *http.ServeMux {
	mux := http.NewServeMux()

	for key, value := range stories {
		if key == "intro" {
			key = ""
		}

		mux.HandleFunc(fmt.Sprintf("/%s", key), storyHandler(value, templ))
	}

	return mux
}

func main() {

	data, err := os.ReadFile("Story/story.json")
	check(err)

	// Create map that will hold all stories.
	stories := make(map[string]Arc, 0)

	// Unmarshall the json and store the values in the map.
	err = json.Unmarshal(data, &stories)
	check(err)

	// Establishes the template that will execute
	tmpl := template.Must(template.ParseFiles("Templates/story.html"))

	mux := defaultMux(stories, tmpl)

	http.ListenAndServe(":8000", mux)
}
