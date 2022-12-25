package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
description: error handling
param: err error
*/
func errorHanding(err error) {
	if err != nil {
		log.Fatalf("Err: ", err)
	}
}

/*
description: convert map to json object
param: statesAndStateCapitalsInMap map[string]string
return: statesAndStateCapitalsInJson []byte
*/
func mapToJson(statesAndStateCapitalsInMap map[string]string) []byte {

	statesAndStateCapitalsInJson, err := json.MarshalIndent(statesAndStateCapitalsInMap, "", "  ")
	errorHanding(err)
	return statesAndStateCapitalsInJson
}

func main() {

	statesAndStateCapitalsInMap := map[string]string{
		"Andhra Pradesh":    "Amaravati",
		"Arunachal Pradesh": "Itanagar",
		"Assam":             "Dispur",
		"Bihar":             "Patna",
		"Maharashtra":       "Mumbai",
		"Jharkhand":         "Ranchi",
		"Karnataka":         "Bangalore",
		"Kerala":            "Thiruvananthapuram",
		"Madhya Pradesh":    "Bhopal",
		"Haryana":           "Chandigarh"}

	stateAndStateCapitalsInJson := string(mapToJson(statesAndStateCapitalsInMap))
	fmt.Println(stateAndStateCapitalsInJson)
}
