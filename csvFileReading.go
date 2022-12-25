package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

/*
description: to handle error
param: err error
*/
func errorHandling(err error) {
	if err != nil {
		log.Fatalf("Err: ", err)
	}
}

/*
description: Read all the file records with header
param: file *os.File
*/
func readCompleteFile(file *os.File) {

	//print all the record with header file
	records, err := csv.NewReader(file).ReadAll()
	errorHandling(err)
	fmt.Println("Print all the record of file with header line")
	for _, record := range records {
		fmt.Println(record)
	}
	fmt.Println("")
}

/*
description: Read all the file records without header file
param: file *os.File
*/
func readCompleteFileExceptHeaderLine(file *os.File) {

	//print all the records without header file
	records2 := csv.NewReader(file)
	records2.Read()
	allRecords, err := records2.ReadAll()
	errorHandling(err)
	fmt.Println("Print all the records of file without header line")
	for _, record := range allRecords {
		fmt.Println(record)
	}
}

func main() {

	file, err := os.Open("Cricket_Players_Stats.csv")
	errorHandling(err)
	readCompleteFile(file)

	file, err = os.Open("Cricket_Players_Stats.csv")
	errorHandling(err)
	defer file.Close()
	readCompleteFileExceptHeaderLine(file)
}
