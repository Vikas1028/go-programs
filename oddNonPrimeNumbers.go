package main

import (
	"fmt"
)

/*
description: Input number is a nonprime number then return true and

	if number is a prime number then return false

param: belowNumbersOfInputNum int
*/
func nonPrimeNumber(belowNumberOfInputNum int) bool {
	countOfDivisorsOfBelowNumbersOfInputNum := 0

	for divisors := 2; divisors < belowNumberOfInputNum; divisors++ {
		if belowNumberOfInputNum%divisors == 0 {
			countOfDivisorsOfBelowNumbersOfInputNum++
		}
	}

	if countOfDivisorsOfBelowNumbersOfInputNum >= 1 {
		return true
	} else {
		return false
	}

}

/*
discription: compute non prime number is even or odd and print the result true or false
param: nonPrimeNumber int
*/
func oddNonPrimeNumber(nonPrimeNumber int) int {

	if nonPrimeNumber%2 != 0 {
		return 1
	} else {
		return 0
	}
}

func main() {

	var inputNum int
	countOddNonPrimeNumbers := 1

	fmt.Print("Enter any number: ")
	fmt.Scanln(&inputNum)

	for belowNumbersOfInputNum := 2; belowNumbersOfInputNum < inputNum; belowNumbersOfInputNum++ {
		if nonPrimeNumber(belowNumbersOfInputNum) == true {
			countOddNonPrimeNumbers += oddNonPrimeNumber(belowNumbersOfInputNum)
		}
	}

	fmt.Println("Non Prime odd Numbers count: ", countOddNonPrimeNumbers)
}
