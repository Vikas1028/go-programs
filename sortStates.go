package main

import (
	"fmt"
	"sort"
)

func main() {

	stateCapital := map[string]string{}
	stateCapital["Kerala"] = "Thiruvananthapuram"
	stateCapital["Assam"] = "Dispur"
	stateCapital["Maharashtra"] = "Mumbai"
	stateCapital["Jharkhand"] = "Ranchi"
	stateCapital["Andhra Pradesh"] = "Amaravati"
	stateCapital["Karnataka"] = "Bangalore"
	stateCapital["Madhya Pradesh"] = "Bhopal"
	stateCapital["Arunachal Pradesh"] = "Itanagar"
	stateCapital["Haryana"] = "Chandigarh"
	stateCapital["Bihar"] = "Patna"

	states := []string{}
	for key, _ := range stateCapital {
		states = append(states, key)
	}

	sort.Strings(states)
	fmt.Println(states)
}
