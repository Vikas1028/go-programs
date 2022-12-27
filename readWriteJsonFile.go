package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type cricketPlayer struct {
	Name          string  `json:"Name"`
	Period        string  `json:"Period"`
	MatchesPlayed int     `json:"MatchesPlayed"`
	RunsScored    int     `json:"RunsScored"`
	AvgScore      float32 `json:"AvgScore"`
}

/*
description: If error occurs print the error and terminate the program
param: err error
*/
func errorHandling(err error) {
	if err != nil {
		log.Fatalf("Err: ", err)
	}
}

/*
description: Read json file data by adding file data of cricket players in slice of stuct
return: file data
*/
func readFile() []cricketPlayer {

	file, err := os.Open("players.json")
	errorHandling(err)
	defer file.Close()

	var cricketPlayers []cricketPlayer
	data := json.NewDecoder(file)
	err = data.Decode(&cricketPlayers)
	return cricketPlayers
}

/*
description: Update file data and write in another json file
param: cricketPlayers []cricketPlayer (cricket players data of file)
*/
func writeFile(cricketPlayers []cricketPlayer) {

	for index, data := range cricketPlayers {
		data.MatchesPlayed += 7
		data.RunsScored += 300
		cricketPlayers[index] = data
	}

	newFile, err := os.Create("UpdatePlayersData.json")
	errorHandling(err)
	defer newFile.Close()

	jsonData, err := json.MarshalIndent(cricketPlayers, "", "  ")
	_, err = newFile.Write(jsonData)
	fmt.Println(string(jsonData))
}

func main() {

	fmt.Println(readFile())

	writeFile(readFile())
}
