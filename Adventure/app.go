package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
)

type Message struct {
	Message string
	Code    int
}

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

func main() {
	data, err := os.ReadFile("Story/story.json")
	check(err)

	// Create map that will hold all stories.
	stories := make(map[string]Arc, 0)

	// Unmarshall the json and store the values in the map.
	err = json.Unmarshal(data, &stories)
	check(err)

	tmpl := template.Must(template.ParseFiles("Templates/story.html"))

	for key, _ := range stories {
		http.HandleFunc("/"+key, func(rw http.ResponseWriter, r *http.Request) {
			tmpl.Execute(rw, stories[key])
		})
	}

	http.ListenAndServe(":8000", nil)
}
