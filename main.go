package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func ParseCSV(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Cannot read input file "+fileName, err)
	}
	csvReader := csv.NewReader(file)
	csvContent, readerErr := csvReader.ReadAll()
	for _, problem := range csvContent {
		problem[0] = strings.TrimSpace(problem[0])
		problem[1] = strings.TrimSpace(problem[1])
	}
	if readerErr != nil {
		log.Fatal("Cannot parse file "+fileName+" as CSV", readerErr)
	}

	err = file.Close()
	if err != nil {
		log.Println("Re-closing the file ", file.Name())
	}
	return csvContent
}

func askInput() {
	fmt.Print("Press Enter to start quiz.")
	fmt.Scanln()
}

func main() {
	quizFile := flag.String("csv", "problems.csv",
		"a csv file in the format of \"question, answer\" (default \"problems.csv\"")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds (default 30)")
	flag.Parse()
	quiz := ParseCSV(*quizFile)
	askInput()
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	var correct uint
	for i, curLine := range quiz {
		fmt.Print("Problem #", i+1, ": ", curLine[0], " = ")
		answerCh := make(chan string)
		go func() {
			var userAnswer string
			fmt.Scanf("%s\n", &userAnswer)
			answerCh <- userAnswer
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nTime is up. You scored %d out of %d.\n", correct, len(quiz))
			return
		case answer := <-answerCh:
			if answer == curLine[1] {
				correct++
			}
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(quiz))
}
