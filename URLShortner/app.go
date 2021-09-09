package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/GibranHL0/Gophercises/URLShortner/URLShort"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hey there! 🐙")
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func main() {
	mux := defaultMux()
	port := os.Getenv("PORT")

	// Build the MapHandler using the mux as the fallback
	pathToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := urlshort.MapHandler(pathToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the fallback
	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)

	if err != nil {
		panic(err)
	}

	log.Fatal(http.ListenAndServe(":"+port, yamlHandler))
}
