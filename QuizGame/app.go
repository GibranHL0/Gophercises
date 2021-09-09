package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type Problem struct {
	Question string
	Answer   string
}

func getFlags() (path string, timeLimit uint) {
	flag.StringVar(&path, "path", "problems.csv", "Defines the CSV location")
	flag.UintVar(&timeLimit, "time", 5, "Defines the time to complete the test")

	flag.Parse()

	return path, timeLimit
}

func getCSVInfo(path string) [][]string {
	csvFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	csvLines, err := csv.NewReader(csvFile).ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return csvLines
}

func runTest(csvLines [][]string, timeLimit uint) (int, int) {
	correct := 0
	questions := len(csvLines)

	var ready string
	fmt.Println("Press any letter to start, or q to QUIT")
	fmt.Scanln(&ready)

	if ready == "q" || ready == "Q" {
		os.Exit(0)
	}

	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	for index, line := range csvLines {
		problem := Problem{
			Question: line[0],
			Answer:   line[1],
		}
		answerCh := make(chan string)

		fmt.Printf("Problem #%d: %s = ", index+1, problem.Question)

		go func() {
			var userAnswer string
			fmt.Scanln(&userAnswer)
			answerCh <- userAnswer
		}()

		select {
		case <-timer.C:
			return correct, questions
		case userAnswer := <-answerCh:
			if userAnswer == "" {
				return correct, questions
			}

			if userAnswer == problem.Answer {
				correct++
			}
		}
	}

	return correct, questions
}

func main() {
	path, timeLimit := getFlags()
	csvLines := getCSVInfo(path)

	correctAnswers, questions := runTest(csvLines, timeLimit)
	fmt.Printf("\nYou scored %d out of %d!\n", correctAnswers, questions)
}
