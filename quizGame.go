package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
description:  If error is not nil then print error and terminate the program
param: err error
*/
func errorHandling(err error) {
	if err != nil {
		log.Fatalf("Err: ", err)
	}
}

/*
description: read one Quiz CSV file
return: quizQuestion [][]string
*/
func readQuizFile() [][]string {

	csvFile, err := os.Open("quiz.csv")
	errorHandling(err)
	defer csvFile.Close()

	readcsvFile := csv.NewReader(csvFile)
	quizQuestions, err := readcsvFile.ReadAll()
	errorHandling(err)

	return quizQuestions
}

/*
description: play a quiz game
param: quizQuestions [][]string
*/
func quizGame(quizQuestions [][]string) {

	fmt.Println("Welcome to the quiz game")

	for index, quizData := range quizQuestions {

		fmt.Printf("Question %v: %v\n", index+1, string(quizData[0]))
		fmt.Printf("A) %v\n", string(quizData[1]))
		fmt.Printf("B) %v\n", string(quizData[2]))
		fmt.Printf("C) %v\n", string(quizData[3]))
		fmt.Printf("D) %v\n", string(quizData[4]))
		fmt.Print("Enter your selection: ")

		var UserSelection string
		fmt.Scanln(&UserSelection)

		if strings.ToUpper(UserSelection) == string(quizData[5]) {
			fmt.Println("Correct\n")
		} else {
			fmt.Printf("Incorrect :(Correct answer is %v)\n\n", string(quizData[5]))
		}
	}
	fmt.Println("Quiz complete")
}

func main() {

	quizQuestions := readQuizFile()
	quizGame(quizQuestions)
}
