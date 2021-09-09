package urlshort

import (
	"net/http"
	"encoding/json"

	"gopkg.in/yaml.v2"
)

type pathUrl struct {
	Path string `yaml:"path" json:"path"`
	Url  string `yaml:"url" json:"url"`
}

// Converts []byte JSON format into a slice of pathURL structs
func jsonTopathUrl(JSON []byte) ([]pathUrl, error){
	var pathUrls []pathUrl

	err := json.Unmarshal(JSON, &pathUrls)

	if err != nil {
		return nil, err
	}

	return pathUrls, nil
}

// Converts []byte YAML format into a slice of pathUrl structs
func yamlTopathUrl(yml []byte) ([]pathUrl, error){
	var pathUrls []pathUrl

	err := yaml.Unmarshal(yml, &pathUrls)

	if err != nil {
		return nil, err
	}

	return pathUrls, nil
}

// Converts []pathUrl structs into a map[string]string
func pathUrlToMap(pathUrls []pathUrl) map[string]string {
	pathToUrls := make(map[string]string)

	for _, pu := range pathUrls {
		pathToUrls[pu.Path] = pu.Url
	}

	return pathToUrls
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}

		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	pathUrls, err := yamlTopathUrl(yml)

	if err != nil {
		return nil, err
	}

	pathToUrls := pathUrlToMap(pathUrls)

	return MapHandler(pathToUrls, fallback), nil
}


// JSONHandler will parse the provided JSON and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the JSON, then the
// fallback http.Handler will be called instead.
func JSONHandler(JSON []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := jsonTopathUrl(JSON)
	
	if err != nil {
		return nil, err
	}

	pathToUrls := pathUrlToMap(pathUrls)

	return MapHandler(pathToUrls, fallback), nil
}