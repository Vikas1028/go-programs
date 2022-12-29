package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Friend struct {
	Name              string `json:"Name"`
	MobileNo          int    `json:"MobileNo"`
	AlternateMobileNo int    `json:"Alternate MobileNo"`
	Address           string `json:"Address"`
	City              string `json:"City"`
}

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
description: Read json file convert data into slice and print
*/
func readFile() {

	readFile, err := ioutil.ReadFile("addressBook.json")
	errorHandling(err)

	var Friends = []Friend{}

	err = json.Unmarshal(readFile, &Friends)

	for _, data := range Friends {
		fmt.Println(data)
	}
}

func main() {

	readFile()
}
