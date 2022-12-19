package main

import (
	"fmt"
)

func main() {
	var choice int
	fmt.Println("0) Centigrade to Fahrenheit \n1) Fahrenheit to Centigrade")
	fmt.Print("Enter your Choice: ")
	fmt.Scanln(&choice)
	tempratureConverter(choice)
}

func tempratureConverter(choice int) int {
	var temprature float64
	if choice == 0 {
		fmt.Println("Enter Temprature in Centigrade: ")
		fmt.Scanln(&temprature)
		fmt.Println("Temprature in Fahrenheit: ", (temprature*9.0/5.0)+32.0, " degree F.")
	} else if choice == 1 {
		fmt.Print("Enter Temprature in Fahrenheit: ")
		fmt.Scanln(&temprature)
		fmt.Println("Temprature in Centigrade: ", (temprature-32.0)*5.0/9.0, " degree C.")
	} else {
		fmt.Println("Entered a Wrong Choice")
	}
	return 0
}
