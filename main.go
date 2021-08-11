package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func ParseCSV(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Cannot read input file "+fileName, err)
	}
	csvReader := csv.NewReader(file)
	csvContent, readerErr := csvReader.ReadAll()
	if readerErr != nil {
		log.Fatal("Cannot parse file "+fileName+" as CSV", readerErr)
	}

	err = file.Close()
	if err != nil {
		log.Println("Re-closing the file ", file.Name())
	}
	return csvContent
}

func main() {
	//todo: add a timer based on timeLimit flag:
	quizFile := flag.String("csv", "problems.csv",
		"a csv file in the format of \"question, answer\" (default \"problems.csv\"")
	timeLimit := flag.Int("limit", 10, "the time limit for the quiz in seconds (default 10)")
	flag.Parse()
	var correct uint
	fmt.Println("Time limit for this quiz = ", *timeLimit)
	quiz := ParseCSV(*quizFile)
	for i, curLine := range quiz {
		var userAnswer string
		fmt.Print("Problem #", i+1, ": ", curLine[0], " = ")
		fmt.Scanf("%s\n", &userAnswer)
		if userAnswer == curLine[1] {
			correct++
		}
	}
	fmt.Println("You scored ", correct, " out of ", len(quiz), ".")
}
