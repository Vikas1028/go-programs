package main

import (
	"fmt"
)

/*
description: return fibonacci series even numbers count
param: fibonacciSeriesNumber, fibonacciSeriesNumber2, fibonacciSeriesSize, countFibonacciSeriesNumbers, countFibonacciSeriesEvenNumbers int
*/
func fibonacciSeries(fibonacciSeriesNumber, fibonacciSeriesNumber2, fibonacciSeriesSize, countFibonacciSeriesNumbers, countFibonacciSeriesEvenNumbers int) int {

	for countFibonacciSeriesNumbers < fibonacciSeriesSize {
		temp := fibonacciSeriesNumber2
		fibonacciSeriesNumber2 += fibonacciSeriesNumber
		if fibonacciSeriesNumber2%2 == 0 {
			countFibonacciSeriesEvenNumbers++
		}
		fibonacciSeriesNumber = temp
		countFibonacciSeriesNumbers++
	}

	return countFibonacciSeriesEvenNumbers

}

func main() {
	fibonacciSeriesFirstNumber := 0
	fibonacciSeriesSecondNumber := 1
	countFibonacciSeriesNumbers := 2
	countFibonacciSeriesEvenNumbers := 0
	var fibonacciSeriesSize int

	fmt.Print("Enter any Number: ")
	fmt.Scanln(&fibonacciSeriesSize)

	fmt.Println("Fibonacci series even numbers count: ", fibonacciSeries(fibonacciSeriesFirstNumber, fibonacciSeriesSecondNumber, fibonacciSeriesSize, countFibonacciSeriesNumbers, countFibonacciSeriesEvenNumbers))
}
