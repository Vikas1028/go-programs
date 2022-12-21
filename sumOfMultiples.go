package main

import (
	"fmt"
)

/*
description: Sum of all the multiples of 3 or 5 below inputNumber.
param: inputNum
*/
func sumOfMultiples(inputNum int) int {

	sumOfMultiplesOf3And5 := 0

	for belowNumbersOfInputNum := 3; belowNumbersOfInputNum < inputNum; belowNumbersOfInputNum++ {
		if belowNumbersOfInputNum%3 == 0 || belowNumbersOfInputNum%5 == 0 {
			sumOfMultiplesOf3And5 += belowNumbersOfInputNum
		}
	}

	return sumOfMultiplesOf3And5
}

func main() {

	var inputNum int
	fmt.Print("Enter any number: ")
	fmt.Scanln(&inputNum)

	fmt.Println("sum of multiples of 3 or 5: ", sumOfMultiples(inputNum))
}
