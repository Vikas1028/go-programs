package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
description: to handle Error
param: err error
*/
func errorHanding(err error) {
	if err != nil {
		log.Println("Err: ", err)
	}
}

/*
description: convert json object to map structure
param: addressBookInJson []byte
return: addressBookInMap map[string]string
*/
func jsonToMap(addressBookInJson []byte) map[string]string {

	addressBookInMap := map[string]string{}

	err := json.Unmarshal([]byte(addressBookInJson), &addressBookInMap)
	errorHanding(err)

	return addressBookInMap
}

func main() {

	addressBookInJson := `{
		"Thinly Sliced Peeled Apples" : "6 Cups",
		"sugar" : "3/4 cup",
		"flour" : "2 tablespoons",
		"cinamon" : "1/4 teaspoon",
		"nutmeg" : "1/8 tablespoon",
		"lemon juice": "1 tablespoon"}`

	addressBookInMap := jsonToMap([]byte(addressBookInJson))
	for key, value := range addressBookInMap {
		fmt.Println(key, " : ", value)
	}
}
