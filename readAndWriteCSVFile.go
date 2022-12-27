package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func errorHandling(err error) {
	if err != nil {
		log.Fatalf("Err: ", err)
	}
}

func readAndUpdateCSVFile() {

	file, err := os.Open("Cricket_Players_Stats.csv")
	errorHandling(err)
	defer file.Close()

	newFile, err := os.Create("update_Cricket_players_stats.csv")
	errorHandling(err)
	defer newFile.Close()

	readFile := csv.NewReader(file)
	data, err := readFile.Read()
	writeData := fmt.Sprintf("%s,%s,%s,%s,%s\n", data[0], data[1], data[2], data[3], data[4])
	_, err = newFile.WriteString(writeData)
	fmt.Print(writeData)

	for i := 0; i < 5; i++ {
		data, err = readFile.Read()

		matchesPlayed, err := strconv.Atoi(data[2])
		errorHandling(err)
		matchesPlayed += 10

		totalRuns, err := strconv.Atoi(data[3])
		errorHandling(err)
		totalRuns += 350

		writeData := fmt.Sprintf("%s,%s,%d,%d,%s\n", data[0], data[1], matchesPlayed, totalRuns, data[4])
		_, err = newFile.WriteString(writeData)
		fmt.Print(writeData)

	}

}

func main() {

	readAndUpdateCSVFile()
}
