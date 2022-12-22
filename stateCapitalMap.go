package main

import (
	"fmt"
)

func main() {

	stateCapital := map[string]string{}
	stateCapital["Andhra Pradesh"] = "Amaravati"
	stateCapital["Arunachal Pradesh"] = "Itanagar"
	stateCapital["Assam"] = "Dispur"
	stateCapital["Bihar"] = "Patna"
	stateCapital["Maharashtra"] = "Mumbai"
	stateCapital["Jharkhand"] = "Ranchi"
	stateCapital["Karnataka"] = "Bangalore"
	stateCapital["Kerala"] = "Thiruvananthapuram"
	stateCapital["Madhya Pradesh"] = "Bhopal"
	stateCapital["Haryana"] = "Chandigarh"

	fmt.Printf("State Capital Data Type: %T \nStateCapital: %v \n\n", stateCapital, stateCapital)

	fmt.Println("Iterate over map and print state name and its capital name one by one")
	for state, capital := range stateCapital {
		fmt.Println(state, "->", capital)
	}

}
