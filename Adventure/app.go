package main

import (
	"html/template"
	"net/http"

	"github.com/GibranHL0/Gophercises/Adventure/Utils"
	"github.com/GibranHL0/Gophercises/Adventure/Handlers"
)

func main() {
	// Obtain the story
	story := utils.GetStory()

	// Establishes the template that will execute the Handler
	tmpl := template.Must(template.ParseFiles("Templates/story.html"))

	mux := handlers.StoryMux(story, tmpl)

	http.ListenAndServe(":8000", mux)
}
