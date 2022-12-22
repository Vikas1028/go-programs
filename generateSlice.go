package main

import (
	"fmt"
	"strings"
)

func main() {

	var numbersString string
	fmt.Print("Enter sequence of comma-separated numbers: ")
	fmt.Scanln(&numbersString)

	sliceOfStringNumbers := strings.Split(numbersString, ",")
	fmt.Println("Slice of Numbers in String: ", sliceOfStringNumbers)

}
