package main

import (
	"fmt"
)

func main() {

	var inputNum int
	fmt.Print("Enter any Number: ")
	fmt.Scanln(&inputNum)

	if inputNum%3 == 0 {
		fmt.Print("Pling")
	}
	if inputNum%5 == 0 {
		fmt.Print("Plang")
	}
	if inputNum%7 == 0 {
		fmt.Print("Plong")
	}
	if inputNum%3 != 0 && inputNum%5 != 0 && inputNum%7 != 0 {
		fmt.Println(inputNum)
	}
}
