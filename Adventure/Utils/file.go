package utils

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/GibranHL0/Gophercises/Adventure/Story"
)

func getFile() string {
	filepath := flag.String("file", "Files/story.json", "The JSON file with the Adventure story")
	flag.Parse()

	return *filepath
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getFileData(filepath string) []byte {
	data, err := os.ReadFile(filepath)
	check(err)

	return data
}

func GetStory() story.Story{
	// Get the path of the JSON
	filepath := getFile()

	// Obtain the byte slice containing all the information
	data := getFileData(filepath)

	// Create the map that will hold all chapters
	story := story.Story{}

	// Unmarshall the json and store the values in story
	err := json.Unmarshal(data, &story)
	check(err)

	return story
}