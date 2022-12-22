package main

import (
	"fmt"
)

/*
Description: Process PlacesNames array and generate array of placesNamesLength and
store length of String in each element and return placesNamesLength array
param: placesNames
*/
func namesLength(placesNames [7]string) [7]int {
	var placesNamesLength [7]int
	for i := 0; i < len(placesNames); i++ {
		placesNamesLength[i] = len(placesNames[i])
	}
	return placesNamesLength
}

func main() {
	fmt.Print("Enter 7 places' names: ")
	var placesNames [7]string
	fmt.Scanln(&placesNames[0], &placesNames[1], &placesNames[2], &placesNames[3], &placesNames[4], &placesNames[5], &placesNames[6])

	fmt.Println(namesLength(placesNames))
}
