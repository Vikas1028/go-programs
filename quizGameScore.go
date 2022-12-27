package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type quizQuestion struct {
	Question string
	OptionA  string
	OptionB  string
	OptionC  string
	OptionD  string
	Answer   string
}

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
	fileData, err := readcsvFile.ReadAll()
	errorHandling(err)

	return fileData
}

/*
description: Convert file data into slice of struct (slice of struct contain quiz questions)
param: fileData [][]string
return: quizQuestions []quizQuestion
*/
func quizQuestions(fileData [][]string) []quizQuestion {

	var quizQuestions []quizQuestion

	for _, quizData := range fileData {

		quizQuestions = append(quizQuestions, quizQuestion{
			Question: quizData[0],
			OptionA:  quizData[1],
			OptionB:  quizData[2],
			OptionC:  quizData[3],
			OptionD:  quizData[4],
			Answer:   quizData[5]})
	}

	return quizQuestions
}

/*
description: Start quize and ask five random questions
param: quizQuestion []quizQuestion
return: score
*/
func startQuize(quizQuestions []quizQuestion) int {

	score := 0
	sliceOfAskedQuestions := []int{}
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Welcome to the quiz game")

	for len(sliceOfAskedQuestions) < 5 {

		questionNo := rand.Intn(10) + 1
		questionPresent := false
		for _, num := range sliceOfAskedQuestions {
			if num == questionNo {
				questionPresent = true
			}
		}

		if questionPresent == false {
			sliceOfAskedQuestions = append(sliceOfAskedQuestions, questionNo)
			score = quizGame(quizQuestions, questionNo, score, len(sliceOfAskedQuestions))
		}

	}
	fmt.Println("Quiz complete")
	return score
}

/*
description: Show question on terminal and get answer from user the
param: quizQuestions []quizQuestion, questionNo, score, totalquestionsAsk int
return: score
*/
func quizGame(quizQuestions []quizQuestion, questionNo, score, totalquestionsAsk int) int {

	fmt.Printf("Question %v: %v\n", totalquestionsAsk, quizQuestions[questionNo].Question)
	fmt.Printf("A) %v\n", quizQuestions[questionNo].OptionA)
	fmt.Printf("B) %v\n", quizQuestions[questionNo].OptionB)
	fmt.Printf("C) %v\n", quizQuestions[questionNo].OptionC)
	fmt.Printf("D) %v\n", quizQuestions[questionNo].OptionD)
	fmt.Print("Enter your selection: ")

	var UserSelection string
	fmt.Scanln(&UserSelection)

	if strings.ToUpper(UserSelection) == quizQuestions[questionNo].Answer {
		fmt.Println("Correct\n")
		score++
	} else {
		fmt.Printf("Incorrect :(Correct answer is %v)\n\n", quizQuestions[questionNo].Answer)
	}
	return score
}

func main() {

	fileData := readQuizFile()

	quizeQuestions := quizQuestions(fileData)

	fmt.Println("Total Score : ", startQuize(quizeQuestions))
}
