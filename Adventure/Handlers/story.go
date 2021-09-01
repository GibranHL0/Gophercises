package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/GibranHL0/Gophercises/Adventure/Story"
)

func StoryMux(story story.Story, templ *template.Template) *http.ServeMux {
	mux := http.NewServeMux()

	for key, value := range story {
		if key == "intro" {
			key = ""
		}

		mux.HandleFunc(fmt.Sprintf("/%s", key), storyHandler(value, templ))
	}

	return mux
}

func storyHandler(chapter story.Chapter, tmpl *template.Template) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		tmpl.Execute(rw, chapter)
	}
}