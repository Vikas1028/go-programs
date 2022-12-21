package main

import (
	"fmt"
)

/*
discription: print the fibonacci series
param: fibonacciSeriesNumber, fibonacciSeriesNumber2, fibonacciSeriesSize, countFibonacciSeriesNumbers int
*/
func fibonacciSeries(fibonacciSeriesNumber, fibonacciSeriesNumber2, fibonacciSeriesSize, countFibonacciSeriesNumbers int) {
	countFibonacciSeriesNumbers++

	if countFibonacciSeriesNumbers == fibonacciSeriesSize {
		fmt.Print(fibonacciSeriesNumber + fibonacciSeriesNumber2)
	}

	if countFibonacciSeriesNumbers < fibonacciSeriesSize {
		temp := fibonacciSeriesNumber2
		fmt.Print(fibonacciSeriesNumber+fibonacciSeriesNumber2, ",")
		fibonacciSeriesNumber2 = fibonacciSeriesNumber + fibonacciSeriesNumber2
		fibonacciSeriesNumber = temp
		fibonacciSeries(fibonacciSeriesNumber, fibonacciSeriesNumber2, fibonacciSeriesSize, countFibonacciSeriesNumbers)
	}

}

func main() {
	fibonacciSeriesFirstNumber := 0
	fibonacciSeriesSecondNumber := 1
	countFibonacciSeriesNumbers := 2
	var fibonacciSeriesSize int

	fmt.Print("Enter any Number: ")
	fmt.Scanln(&fibonacciSeriesSize)
	fmt.Print(fibonacciSeriesFirstNumber, ",", fibonacciSeriesSecondNumber, ",")

	fibonacciSeries(fibonacciSeriesFirstNumber, fibonacciSeriesSecondNumber, fibonacciSeriesSize, countFibonacciSeriesNumbers)
}
